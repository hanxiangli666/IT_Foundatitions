package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"smartyihui-web/backend/internal/config"
	"smartyihui-web/backend/internal/store"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	db, err := openDB(cfg)
	if err != nil {
		log.Fatalf("open database failed: %v", err)
	}
	defer db.Close()

	if err := store.Migrate(db); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	messageStore := store.NewMessageStore(db)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.FrontendOrigin,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"ok":   true,
			"time": time.Now().Format(time.RFC3339),
		})
	})

	app.Get("/api/messages", func(c *fiber.Ctx) error {
		messages, err := messageStore.List(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"items": messages})
	})

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var req struct {
			Content string `json:"content"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json body"})
		}
		if req.Content == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "content is required"})
		}

		item, err := messageStore.Create(c.Context(), req.Content)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(item)
	})

	frontendDist := filepath.Clean("../frontend/dist")
	if _, err := os.Stat(filepath.Join(frontendDist, "index.html")); err == nil {
		app.Static("/", frontendDist)
		app.Get("*", func(c *fiber.Ctx) error {
			return c.SendFile(filepath.Join(frontendDist, "index.html"))
		})
	}

	addr := ":" + cfg.AppPort
	log.Printf("server listening on %s", addr)
	log.Fatal(app.Listen(addr))
}

func openDB(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		cfg.MySQLUser,
		cfg.MySQLPassword,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDatabase,
		cfg.MySQLParams,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if maxOpen := os.Getenv("MYSQL_MAX_OPEN_CONNS"); maxOpen != "" {
		// keep compatibility for future tuning without failing startup
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

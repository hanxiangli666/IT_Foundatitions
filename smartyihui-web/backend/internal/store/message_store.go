package store

import (
	"context"
	"database/sql"
	"time"
)

type Message struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type MessageStore struct {
	db *sql.DB
}

func NewMessageStore(db *sql.DB) *MessageStore {
	return &MessageStore{db: db}
}

func (s *MessageStore) List(ctx context.Context) ([]Message, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT id, content, created_at
		FROM messages
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Message, 0)
	for rows.Next() {
		var item Message
		var createdAt time.Time
		if err := rows.Scan(&item.ID, &item.Content, &createdAt); err != nil {
			return nil, err
		}
		item.CreatedAt = createdAt.Format(time.RFC3339)
		items = append(items, item)
	}

	return items, rows.Err()
}

func (s *MessageStore) Create(ctx context.Context, content string) (Message, error) {
	result, err := s.db.ExecContext(ctx, `
		INSERT INTO messages (content)
		VALUES (?)
	`, content)
	if err != nil {
		return Message{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Message{}, err
	}

	var item Message
	var createdAt time.Time
	if err := s.db.QueryRowContext(ctx, `
		SELECT id, content, created_at
		FROM messages
		WHERE id = ?
	`, id).Scan(&item.ID, &item.Content, &createdAt); err != nil {
		return Message{}, err
	}
	item.CreatedAt = createdAt.Format(time.RFC3339)

	return item, nil
}

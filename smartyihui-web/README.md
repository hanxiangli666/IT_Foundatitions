# smartyihui-web

A full-stack web project for building a company portal website:

- **Frontend:** Vue 3 + Vue Router + Naive UI + Axios + Vite
- **Backend:** Go + Fiber
- **Database:** MySQL

## Structure

```text
smartyihui-web/
├─ frontend/
└─ backend/
```

## Frontend setup

```bash
cd frontend
npm install
npm run dev
```

Default dev server: `http://127.0.0.1:31098`

## Backend setup

1. Create a MySQL database, for example:

```sql
CREATE DATABASE smartyihui DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. Copy env file and update it:

```bash
cd backend
cp .env.example .env
```

3. Run the backend:

```bash
go mod tidy
go run .
```

Default API server: `http://127.0.0.1:31098`

## Docker deployment

This repo supports root-level Docker deployment for the frontend and backend together.

Build the image from the repository root:

```bash
docker build -t smartyihui-web .
```

Run the container:

```bash
docker run --rm -p 31098:31098 --env-file backend/.env smartyihui-web
```

After startup, the container listens on the port configured by `APP_PORT`. If you set `APP_PORT=31098`, both the API and built frontend are served on `http://127.0.0.1:31098`.

## Environment variables

`backend/.env.example`

```env
APP_PORT=31098
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=123456
MYSQL_DATABASE=smartyihui
MYSQL_PARAMS=charset=utf8mb4&parseTime=True&loc=Local
FRONTEND_ORIGIN=http://127.0.0.1:31098
```

## Notes

- This repo was initialized on branch `init`.
- The current codebase can be extended into a company portal with modules such as company profile, news, announcements, contact info, and internal navigation.
- Docker expects your MySQL instance to be reachable from inside the container through the values in `backend/.env`.

FROM node:20 AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend ./
RUN npm run build

FROM golang:1.25.3 AS backend-builder
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend ./
RUN go build -o app main.go

FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y ca-certificates tzdata && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app/backend

COPY --from=backend-builder /app/backend /app/backend
COPY --from=frontend-builder /app/frontend/dist /app/frontend/dist

EXPOSE 31098

CMD ["sh", "-c", "cd /app/backend && ./app"]

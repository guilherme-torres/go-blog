-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" INTEGER PRIMARY KEY,
    "name" TEXT NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "password_hash" TEXT NOT NULL,
    "role" TEXT NOT NULL CHECK("role" in ('user', 'editor', 'admin')) DEFAULT 'user',
    "created_at" TEXT DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TEXT DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd

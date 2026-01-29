-- +goose Up
-- +goose StatementBegin
CREATE TABLE "articles" (
    "id" INTEGER PRIMARY KEY,
    "title" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "author_id" INTEGER NOT NULL,
    "published_at" TEXT DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("author_id") REFERENCES "users" ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "articles";
-- +goose StatementEnd

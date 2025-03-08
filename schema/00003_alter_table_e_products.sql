-- +goose Up
-- +goose StatementBegin
ALTER TABLE "products" ADD COLUMN "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "products" DROP COLUMN "created_at";
-- +goose StatementEnd

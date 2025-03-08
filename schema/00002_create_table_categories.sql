-- +goose Up
-- +goose StatementBegin
CREATE TABLE "categories"(
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "name" jsonb NOT NULL,
    "photo" VARCHAR(255)
);

CREATE TABLE "products"(
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "category_id" BIGINT NOT NULL REFERENCES "categories"("id"),
    "name" jsonb NOT NULL,
    "code" VARCHAR(64) NOT NULL,
    "price" BIGINT NOT NULL,
    "photo" VARCHAR(64)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "products";
DROP TABLE "categories";
-- +goose StatementEnd

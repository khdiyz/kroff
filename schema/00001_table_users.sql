-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users"(
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "phone_number" VARCHAR(64) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd

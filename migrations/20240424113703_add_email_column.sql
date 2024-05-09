-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD column email varchar(255) NOT NULL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP column email
-- +goose StatementEnd

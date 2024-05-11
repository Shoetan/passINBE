-- +goose Up
-- +goose StatementBegin
ALTER TABLE vault
ADD column record_email varchar(255) NOT NULL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE vault
DROP column record_email
-- +goose StatementEnd

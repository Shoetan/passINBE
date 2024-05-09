-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  user_id  bigserial primary key, 
  name varchar(255) not null,
  master_password varchar(255) not null,
  inserted_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd

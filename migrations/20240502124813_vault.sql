-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vault (
    record_id bigserial primary key,
    user_id int,
    record_name varchar(255),
    record_password bytea,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vault;
-- +goose StatementEnd

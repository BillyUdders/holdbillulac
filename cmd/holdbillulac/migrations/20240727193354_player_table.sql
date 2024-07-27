-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS players
(
    id   integer primary key,
    name text,
    age  integer,
    MMR  integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE players
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS players
(
    id   integer primary key,
    name text    not null,
    age  integer not null,
    MMR  integer not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE players
-- +goose StatementEnd

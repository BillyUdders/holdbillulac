-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS nav
(
    id       integer primary key,
    name     TEXT,
    nav_data JSONB
);

INSERT INTO nav (name, nav_data)
VALUES ('Home', '{
  "Google": "https://www.google.com",
  "YouTube": "https://www.youtube.com",
  "Facebook": "https://www.facebook.com",
  "Twitter": "https://www.twitter.com"
}'),
       ('Contact', '{
         "LinkedIn": "https://www.linkedin.com",
         "GitHub": "https://www.github.com",
         "Email": "mailto:contact@example.com",
         "Phone": "tel:+123456789"
       }'),
       ('Resources', '{
         "Documentation": "https://docs.example.com",
         "API Reference": "https://api.example.com",
         "Support": "https://support.example.com",
         "Community": "https://community.example.com"
       }');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE player;
-- +goose StatementEnd

-- +goose Up
CREATE TABLE users (
    id int GENERATED ALWAYS AS IDENTITY NOT NULL,
    name varchar(32) NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE users;
-- +migrate Up
CREATE TABLE if not exists todos (
    name text
);

-- +migrate Down
DROP TABLE todos;

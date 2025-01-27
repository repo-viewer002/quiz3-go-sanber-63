-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id          SERIAL        PRIMARY KEY,
  username    varchar(255)  UNIQUE NOT NULL,
  password    varchar(255)  NOT NULL,
  created_at  TIMESTAMP     DEFAULT CURRENT_TIMESTAMP NOT NULL,
  created_by  varchar(255)  NOT NULL,
  modified_at TIMESTAMP     DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_by varchar(255)  NOT NULL
);

-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin

CREATE OR REPLACE FUNCTION update_modified_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER users_modified_at_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_modified_at();

-- +migrate StatementEnd
-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
  id          SERIAL        PRIMARY KEY,
  name        varchar(255)  NOT NULL,
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

CREATE TRIGGER categories_modified_at_trigger
BEFORE UPDATE ON categories
FOR EACH ROW
EXECUTE FUNCTION update_modified_at();

-- +migrate StatementEnd
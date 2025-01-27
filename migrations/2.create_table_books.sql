-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books (
  id            SERIAL        PRIMARY KEY,
  title         varchar(255)  NOT NULL,
  description   varchar(255)  NOT NULL,
  image_url     varchar(255)  NOT NULL,
  release_year  SMALLINT      NOT NULL,
  price         INTEGER       NOT NULL,
  total_page    INTEGER       NOT NULL,
  thickness     varchar(255)  NOT NULL,
  category_id   INTEGER       NOT NULL,
  created_at    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP NOT NULL,
  created_by    varchar(255)  NOT NULL,
  modified_at   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_by   varchar(255)  NOT NULL,

  CONSTRAINT fk_category
    FOREIGN KEY (category_id)
    REFERENCES categories(id)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
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

CREATE TRIGGER books_modified_at_trigger
BEFORE UPDATE ON books
FOR EACH ROW
EXECUTE FUNCTION update_modified_at();

-- +migrate StatementEnd

CREATE TABLE amis (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL
);

CREATE TRIGGER amis_updated_at
  BEFORE UPDATE OR INSERT ON amis
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE scans ADD COLUMN ami_id int REFERENCES amis(id) ON DELETE SET NULL;

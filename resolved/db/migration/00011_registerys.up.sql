
CREATE TABLE registerys (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  vault text DEFAULT NULL,
  vault_path text DEFAULT NULL
);

CREATE TRIGGER registerys_updated_at
  BEFORE UPDATE OR INSERT ON registerys
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();


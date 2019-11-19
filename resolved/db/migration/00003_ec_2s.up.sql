
CREATE TABLE ec_2s (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  ip text DEFAULT NULL,
  dns text DEFAULT NULL,
  aws_guid text DEFAULT NULL
);

CREATE TRIGGER ec_2s_updated_at
  BEFORE UPDATE OR INSERT ON ec_2s
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();


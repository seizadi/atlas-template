
CREATE TABLE knowledge_bases (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL
);

CREATE TRIGGER knowledge_bases_updated_at
  BEFORE UPDATE OR INSERT ON knowledge_bases
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();


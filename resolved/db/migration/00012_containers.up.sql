
CREATE TABLE containers (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  image_tag text DEFAULT NULL,
  digest text DEFAULT NULL,
  registry_id int REFERENCES registrys(id) ON DELETE SET NULL,
  application_id int REFERENCES applications(id) ON DELETE SET NULL
);

CREATE TRIGGER containers_updated_at
  BEFORE UPDATE OR INSERT ON containers
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE scans ADD COLUMN container_id int REFERENCES containers(id) ON DELETE SET NULL;

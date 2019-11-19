
CREATE TABLE scans (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  scan_time timestamptz DEFAULT current_timestamp,
  scan_type text DEFAULT NULL,
  scan_job jsonb NOT NULL DEFAULT '{}',
  state text DEFAULT NULL,
  application_id int DEFAULT NULL REFERENCES applications(id) ON DELETE SET NULL,
  ec2_id int DEFAULT NULL REFERENCES ec2s(id) ON DELETE SET NULL
);

CREATE TRIGGER scans_updated_at
  BEFORE UPDATE OR INSERT ON scans
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();


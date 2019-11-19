
CREATE TABLE vpcs (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  aws_guid text DEFAULT NULL,
  region_id int REFERENCES regions(id) ON DELETE SET NULL
);

CREATE TRIGGER vpcs_updated_at
  BEFORE UPDATE OR INSERT ON vpcs
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE ec2s ADD COLUMN vpc_id int REFERENCES vpcs(id) ON DELETE SET NULL;

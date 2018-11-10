
CREATE TABLE version_tags (
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL
);

CREATE TRIGGER version_tags_updated_at
  BEFORE UPDATE OR INSERT ON version_tags
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();


ALTER TABLE person
DROP COLUMN birthdate;

ALTER TABLE person
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL;

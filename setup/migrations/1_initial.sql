-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE announcements 
(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  title VARCHAR(255),
  content TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY(id)
);

CREATE TABLE profiles
(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  sub VARCHAR(255),
  nickname VARCHAR(255),
  age TINYINT UNSIGNED,
  birthdate DATE,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY(id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE announcements;
DROP TABLE profile;
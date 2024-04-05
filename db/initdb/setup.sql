CREATE DATABASE IF NOT EXISTS household-accounts;
USE household-accounts;

CREATE TABLE IF NOT EXISTS sample
(
  `id`         int(11) NOT NULL AUTO_INCREMENT,
  `name`     text,
  PRIMARY KEY (`id`)
);

INSERT INTO sample (name) VALUES ("sample1");
INSERT INTO sample (name) VALUES ("sample2");
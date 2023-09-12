CREATE DATABASE IF NOT EXISTS lib;

CREATE TABLE `lib`.authors (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name varchar(64) NOT NULL
);

CREATE TABLE `lib`.books (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	title varchar(255) NOT NULL,
  author_id BIGINT,
  
  CONSTRAINT author_id_fk FOREIGN KEY (author_id) REFERENCES authors (id)
);

INSERT INTO `lib`.authors (name) VALUES ('John'), ('Anthony'), ('William');
INSERT INTO `lib`.books (title, author_id) VALUES ('A', 1), ('B', 1), ('C', 2), ('D', 3), ('E', 3), ('F', 3), ('A', 3);

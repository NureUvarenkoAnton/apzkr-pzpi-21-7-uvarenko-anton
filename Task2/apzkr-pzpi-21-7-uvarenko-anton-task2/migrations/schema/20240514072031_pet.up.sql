CREATE TABLE pets (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  owner_id BIGINT REFERENCES users(id),
  name VARCHAR(255),
  age SMALLINT,
  additional_info TEXT
);

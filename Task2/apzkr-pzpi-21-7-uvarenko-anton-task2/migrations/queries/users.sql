-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: CreateUser :exec
INSERT INTO users (name, email, password, user_type) VALUES(?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users
SET 
  name = ?,
  email = ?
WHERE
  id = ?;

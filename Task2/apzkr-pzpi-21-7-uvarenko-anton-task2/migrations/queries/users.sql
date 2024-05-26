-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: GetUserByUserType :many
SELECT * FROM users
WHERE user_type = ?;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (name, email, password, user_type) VALUES(?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users
SET 
  name = ?,
  email = ?
WHERE
  id = ?;

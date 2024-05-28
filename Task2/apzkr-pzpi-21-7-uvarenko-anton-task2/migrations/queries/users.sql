-- name: GetUserById :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: GetUsers :many
SELECT * FROM users
WHERE 
  id = ? OR
  name = ? OR 
  user_type = ? OR
  is_banned = ? OR
  is_deleted = ?;

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

-- name: DeleteUser :exec
UPDATE users
SET is_deleted = true
WHERE id = ?;

-- name: SetBanState :exec
UPDATE users
SET is_banned = ?
WHERE id = ?;

-- name: CreateWalk :exec
INSERT INTO walks
  (owner_id, walker_id, pet_id, start_time, state)
VALUES
  (?, ?, ?, ?, 'pending');

-- name: UpdateWalkState :exec
UPDATE walks
SET 
  state = ?,
  finish_time = ?
WHERE
  id = ?;

-- name: GetWalkInfoByParams :many
SELECT * FROM walk_info
WHERE
  owner_id = sqlc.narg('owner_id') OR
  walker_id = sqlc.narg('walker_id') OR
  pet_id = sqlc.narg('pet_id');

-- name: GetWalkInfoByWalkId :one
SELECT * FROM walk_info
WHERE walk_id = ?;

-- name: GetWalksByWalkerId :many
SELECT * FROM walks
WHERE walker_id = ?;

-- name: GetWalksByOwnerId :many
SELECT * FROM walks
WHERE owner_id = ?;

-- name: GetWalkById :one
SELECT * FROM walks
WHERE id = ?;

-- name: CreateWalk :exec
INSERT INTO walks
  (owner_id, walker_id, pet_id, start_time, state)
VALUES
  (?, ?, ?, ?, 'pending');

-- name: UpdateWalkState :exec
UPDATE walks
SET 
  state = ?
WHERE
  id = ?;

-- name: GetWalksByWalkerId :many
SELECT * FROM walks
WHERE walker_id = ?;

-- name: GetWalksByOwnerId :many
SELECT * FROM walks
WHERE owner_id = ?;

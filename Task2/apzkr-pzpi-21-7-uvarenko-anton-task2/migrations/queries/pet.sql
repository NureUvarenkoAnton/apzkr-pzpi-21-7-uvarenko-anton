-- name: GetAllPetsByOwnerId :many
SELECT * FROM pets
WHERE owner_id = ?;

-- name: GetPetById :one
SELECT * FROM pets
WHERE id = ?;

-- name: AddPet :exec
INSERT INTO pets 
  (owner_id, name, age, additional_info)
VALUES 
  (?, ?, ?, ?);

-- name: UpdatePet :exec
UPDATE pets
SET
  name = ?,
  age = ?,
  additional_info = ?
WHERE
  id = ?;

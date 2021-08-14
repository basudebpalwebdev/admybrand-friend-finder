-- name: CreateUser :one
INSERT INTO users (
	username,
	description,
	dob,
	address,
	created_at
) VALUES (
	$1, $2, $3, $4, $5
) RETURNING *;
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;
-- name: FindUserByUsername :one
SELECT * FROM users
WHERE username = $1;
-- name: ListUsers :many
SELECT 
	username,
	description,
	dob,
	address,
	created_at
FROM users
ORDER BY id
LIMIT $1
OFFSET $2;
-- name: UpdateUserDetails :one
UPDATE users
SET 
	description = $2,
	dob = $3,
	address = $4
WHERE id = $1
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
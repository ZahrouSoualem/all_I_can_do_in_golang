-- name: CreateUser :one
INSERT INTO users (
username,
hash_password
) VALUES ($1,$2) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsersList :many
SELECT * FROM users;

-- name: GetUserByName :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;

-- name: UpdateUser :one
UPDATE users SET username = $2
WHERE id = $1
RETURNING *;




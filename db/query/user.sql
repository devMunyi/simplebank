-- name: CreateUser :one
INSERT INTO users (
    name, email
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1;


-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;


-- name: UpdateUser :one
UPDATE users SET name = $2, email = $3 
WHERE id = $1
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id DESC LIMIT $1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
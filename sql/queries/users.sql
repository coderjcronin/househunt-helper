-- name: createUser :one
INSERT INTO users (id, created_at, updated_at, name, email, hashed_password) 
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
) 
RETURNING *;

-- name: getUserByUUID :one
SELECT * FROM users WHERE id = $1;

-- name: getUserByName :one
SELECT * FROM users WHERE name = $1;

-- name: wipeUserTable :exec
DELETE FROM users;

-- name: listUserTable :many
SELECT name, email FROM users;
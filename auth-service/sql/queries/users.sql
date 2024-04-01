-- name: CreateUser :one
INSERT INTO users (
    id,
    first_name,
    last_name,
    email_address,
    password,
    user_active,
    created_at,
    updated_at 
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM users u WHERE u.email_address = $1;


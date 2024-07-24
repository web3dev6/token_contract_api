-- name: CreateTransaction :one
INSERT INTO transactions (username, context, payload)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetTransaction :one
SELECT *
FROM transactions
WHERE id = $1
LIMIT 1;
-- name: ListTransactions :many
SELECT *
FROM transactions
WHERE username = $1
ORDER BY id
LIMIT $2 OFFSET $3;

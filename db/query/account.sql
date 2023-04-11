-- name: CreateAccount :one
INSERT INTO accounts (
    owner, balance, currency
) VALUES (
             $1, $2, $3
         )
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
where id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
where id = $1 LIMIT 1 for NO KEY UPDATE  ;

-- name: GetAccountBalance :one
SELECT * FROM accounts
where id = $1 LIMIT 1 for NO KEY UPDATE  ;

-- name: GetAccounts :many
SELECT * FROM accounts
order by id
LIMIT $1
OFFSET $2;



-- name: UpdateAccounts :one
UPDATE accounts
set balance = $2
WHERE id = $1
RETURNING *;


-- name: AddAccountBalance :one
UPDATE accounts
set balance = balance + sqlc.arg(amount)
WHERE id =  sqlc.arg(id)
RETURNING *;





-- name: DeleteAccounts :exec
DELETE FROM accounts
WHERE id = $1;




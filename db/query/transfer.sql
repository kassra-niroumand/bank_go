-- "from_account_id" bigint NOT NULL,
--                              "amount" bigint NOT NULL,
--                              "to_account_id" bigint NOT NULL,
-- name: CreateTransfer :one
insert into transfers(
                      from_account_id, to_account_id, amount
                      )
VALUES (
        $1, $2, $3
) returning *;

-- name: GetTransfer :one
select * from transfers
where id = $1 limit 1;

-- name: ListTransfers :many
select * from transfers
where
    from_account_id = $1 OR
    to_account_id = $2
order by id
limit $3
offset $4;


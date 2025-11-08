-- name: ListChilds :many
SELECT * FROM childs;

-- name: CreateChild :execresult
INSERT INTO childs (`name`, `birthday`, `nickname`) VALUES (
    ?, ?, ? 
);

-- name: DeleteChild :exec
DELETE FROM childs WHERE id = ?;
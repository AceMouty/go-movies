-- name: GetUserByEmail :one
SELECT *
FROM public.users as u
WHERE u.email = $1;

-- name: GetUserById :one
SELECT *
FROM public.users u
WHERE u.id = $1;

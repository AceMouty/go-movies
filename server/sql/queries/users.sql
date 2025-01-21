-- name: GetUserByEmail :one
SELECT *
FROM public.users as u
WHERE u.email = $1;

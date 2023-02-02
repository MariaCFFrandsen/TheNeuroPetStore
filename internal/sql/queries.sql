-- name: GetPets :many
SELECT
   *
FROM
    public.pets;

-- name: BuyPet :exec
UPDATE pets SET available = false WHERE id = $1;
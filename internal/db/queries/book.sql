-- name: GetBook :one
SELECT
    b.id,
    b.title,
    b.isbn,
    b.description,
    b.price,
    b.published_date,
    b.created_at,
    b.updated_at,
    a.id as author_id,
    a.name as author_name,
    a.bio as author_bio,
    a.created_at as author_created_at,
    a.updated_at as author_updated_at
FROM books b
         JOIN authors a ON a.id = b.author_id
WHERE b.id = $1;


-- name: ListBooks :many
SELECT
    b.id,
    b.title,
    b.isbn,
    b.description,
    b.price,
    b.published_date,
    b.created_at,
    b.updated_at,
    a.id as author_id,
    a.name as author_name,
    a.bio as author_bio,
    a.created_at as author_created_at,
    a.updated_at as author_updated_at
FROM books b
         JOIN authors a ON a.id = b.author_id
ORDER BY b.created_at DESC;

-- name: CreateBook :one
WITH inserted_book AS (
    INSERT INTO books (
        title,
        isbn,
        description,
        price,
        author_id,
        published_date
    ) VALUES (
        $1, $2, $3, $4, $5, $6
    )
    RETURNING id, title, isbn, description, price, published_date, created_at, updated_at, author_id
)
SELECT
    b.id,
    b.title,
    b.isbn,
    b.description,
    b.price,
    b.published_date,
    b.created_at,
    b.updated_at,
    a.id AS author_id,
    a.name AS author_name,
    a.bio AS author_bio,
    a.created_at AS author_created_at,
    a.updated_at AS author_updated_at
FROM
    inserted_book b
JOIN
    authors a ON a.id = b.author_id;


-- name: UpdateBook :one
UPDATE books
SET
    title = COALESCE($1, title),
    description = COALESCE($2, description),
    price = COALESCE($3, price),
    published_date = COALESCE($4, published_date),
    updated_at = NOW()
WHERE id = $5
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: GetBookByISBN :one
SELECT * FROM books
WHERE isbn = $1;
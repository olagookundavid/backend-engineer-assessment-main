-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: GetAuthorStats :one
SELECT
    a.id AS author_id,
    a.name AS author_name,
    COUNT(b.id) AS total_books,
    COALESCE(AVG(b.price), 0) AS average_book_price,
    COALESCE(MIN(b.published_date)::TEXT, '') AS earliest_publication,
    COALESCE(MAX(b.published_date)::TEXT, '') AS latest_publication,
    COALESCE(SUM(b.price), 0) AS total_revenue,
    jsonb_object_agg(
        EXTRACT(YEAR FROM b.published_date)::TEXT,
        COUNT(b.id)
    ) FILTER (WHERE b.id IS NOT NULL) AS books_by_year
FROM
    authors a
LEFT JOIN
    books b ON a.id = b.author_id
LEFT JOIN
    book_sales s ON b.id = s.book_id
WHERE
    a.id = $1
GROUP BY
    a.id, a.name;

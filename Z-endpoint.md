# API Structure

## Base Path
All endpoints should start with the base path `/api/v1/`.
### Authors
Base Path: `/api/v1/authors`
#### Endpoints:
- `GET /author{id}` - Get Specific authors.
- `GET /authors` - List all authors.
- `GET /authors/{id}/stats` - Get statistics for a specific author.

### Books
Base Path: `/api/v1/books`
#### Endpoints:
- `GET /books` - List all books.
- `POST /books` - Create a new book.
- `GET /books/{id}` - Retrieve details for a specific book.
- `PUT /books/{id}` - Update a book record.
- `DELETE /books/{id}` - Delete a book record.

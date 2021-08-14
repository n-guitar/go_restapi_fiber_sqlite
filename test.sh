# All Books
curl http://localhost:3000/api/v1/book;echo

# Single Book
curl http://localhost:3000/api/v1/book/1;echo

# New Book
curl -X POST http://localhost:3000/api/v1/book;echo

# Delete Book
curl -X DELETE http://localhost:3000/api/v1/book/1;echo

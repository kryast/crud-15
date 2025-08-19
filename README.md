CRUD ke-15

POST
curl -X POST http://localhost:8080/book \
-H "Content-Type: application/json" \
-d '{
  "title": "Belajar Golang Dasar",
  "author": "Ahmad Syarifuddin"
}'


GET
curl http://localhost:8080/book
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
curl http://localhost:8080/book/1

PUT
curl -X PUT http://localhost:8080/book/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Belajar Golang Dasar Updated",
  "author": "Ahmad Syarifuddin Updated"
}'

DELETE curl -X DELETE http://localhost:8080/book/1
Useful commands

### migration
migrate -path=./migrations -database=$GOLANG_PROJECT_DB_DSN up

### GET
curl localhost:4000/v1/printers/{id}

### POST
curl -i -d "$BODY" localhost:4000/v1/printers <br>
BODY='{"name":"Printer1","type":"Laser","is_color":true,"ip_address":"100.101.102.103","status":"Online","supported_paper_sizes":["A4","A3"],"description":"Nice, cheap printer","battery_left":"-1 mins"}'

### PUT
curl -X PUT -d "$BODY" localhost:4000/v1/printers/1/{id} <br>
BODY='{"name":"Test update","type":"fsofsokfs","ip_address":"100.101.102.103","status":"Online","description":"Nice, cheap printer","battery_left":"10000 mins"}'

### DELETE
curl -X DELETE localhost:4000/v1/printers/{id}

### PATCH
curl -X PATCH -d '{"name": "Test PATCH request"}' localhost:4000/v1/printers/1


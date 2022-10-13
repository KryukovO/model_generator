CONF := configs/default.json
LANG := go
OUTPUT := models/

run:
	go run cmd/main.go -config $(CONF) -lang $(LANG) -o $(OUTPUT)

bld:
	go build -o build/mgen cmd/main.go

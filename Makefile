CONF := configs/default.json
LANG := go

run:
	go run cmd/main.go -config $(CONF) -lang $(LANG)
dup:
	docker-compose up
	
ddown:
	docker-compose down

sup:
	go run cmd/orderapp/main.go -path=config/config.yml

script:
	go run scripts/stream-orders/streamorders.go -path=config/config.yml
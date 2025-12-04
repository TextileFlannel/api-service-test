tests:
	go test .\internal\repository\

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down
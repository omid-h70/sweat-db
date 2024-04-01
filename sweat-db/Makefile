#some evans commands
# $show package
# to select pacakge
# $package pb
# $show servic
# to select a service
# $service GreetingServer
# $call sample_service
evans:
	evans --host localhost --port 9090 -r repl

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

clean:
	@echo "Clean and Build"
	docker-compose down -v
	docker-compose build --no-cache
	docker-compose up --force-recreate
	@echo "Clean Build Done!"

down:
	@echo "Stopping docker compose..."
	docker-compose down -v
	@echo "Done!"

tidy:
	go mod tidy

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: proto tidy up clean down evans
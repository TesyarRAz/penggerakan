generate-app-key:
	@echo "Generate App Key"
	@go run cmd/generate_app_key/main.go

compile:
	@echo "Compile Monolith"
	@go build -o dist/bin/monolith cmd/monolith/main.go
	@echo "Compile Migration"
	@go build -o dist/bin/migration_monolith cmd/migration_monolith/main.go
	@echo "Compile generate app key"
	@go build -o dist/bin/generate_app_key cmd/generate_app_key/main.go

serve:
	@echo "Running the server"
	@go run cmd/monolith/main.go

serve-watch:
	@echo "Running the server"
	@air --build.cmd "go build -o dist/bin/monolith/monolith.exe cmd/monolith/main.go" --build.bin "dist\\bin\\monolith\\monolith.exe"

migration:
	@echo "Running migration"
	@go run cmd/migration_monolith/main.go --seed

migration-fresh:
	@echo "Running migration"
	@go run cmd/migration_monolith/main.go --seed --fresh

testing:
	@echo "Running tests"
	@go test -v ./test
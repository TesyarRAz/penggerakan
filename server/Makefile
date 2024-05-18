generate-app-key:
	@echo "Generate App Key"
	@go run cmd/generate_app_key/main.go

compile:
	@echo "Compile Monolith"
	@go build -o dist/bin/monolith/monolith.exe cmd/monolith/main.go

serve:
	@echo "Running the server"
	@air --build.cmd "go build -o dist/bin/monolith/monolith.exe cmd/monolith/main.go" --build.bin "dist\\bin\\monolith\\monolith.exe"

migration:
	@echo "Running migration for user"
	@go run cmd/migration/main.go --seed

testing:
	@echo "Running tests"
	@go test -v ./test
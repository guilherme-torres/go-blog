migrate:
	@echo "🚀 Executando migrations..."
	@goose -dir migrations sqlite3 blog.db up
	@echo "✅ Migrations executadas com sucesso!"

build:
	@echo "🛠️ Iniciando compilação do binário..."
	@go build -o bin/app ./cmd/server
	@echo "✅ Build concluído com sucesso!"

start: build migrate
	@echo "🚀 Iniciando servidor de produção..."
	./bin/app

tidy:
	@go mod tidy

dev: migrate
	@echo "🚀 Iniciando servidor de desenvolvimento..."
	go run ./cmd/server/main.go

clean:
	@rm -rf bin/
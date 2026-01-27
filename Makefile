APP="My daily control questions"

.PHONY:	run_linter run_app

run_linter:
	golangci-lint run

run_app:
	set -a; . .env; set +a; go run cmd/main.go ask

# Запуск приложения с нуля для проверки работы
start_from_scratch:
	rm -rf data
	mkdir data/
	set -a; . .env; set +a; go run cmd/main.go ask

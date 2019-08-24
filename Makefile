test:
	docker-compose exec app reflex -r '(\.go$$|go\.mod)' -s go test ./...
up:
	docker-compose up app

test:
	docker-compose exec app reflex -r '(\.go$$|go\.mod)' -s go test ./...
test-force:
	docker-compose exec app reflex -r '(\.go$$|go\.mod)' -s go test ./... -- -count=1
up:
	docker-compose up app

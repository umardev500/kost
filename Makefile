up:
	migrate -database "postgres://root:root@127.0.0.1:5432/kost?sslmode=disable" -path database/migrations up
down:
	migrate -database "postgres://root:root@127.0.0.1:5432/kost?sslmode=disable" -path database/migrations down
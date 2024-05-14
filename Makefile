migrate-up:
	migrate -path migrations -database "postgresql://postgres:nodirbek@localhost:5432/productdb?sslmode=disable" -verbose up

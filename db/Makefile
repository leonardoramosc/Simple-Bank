migrateup: 
	migrate -path migration -database "postgresql://postgres:12345@localhost:5432/bank?sslmode=disable" -verbose up

migratedown: 
	migrate -path migration -database "postgresql://postgres:12345@localhost:5432/bank?sslmode=disable" -verbose down
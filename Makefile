#!make
include config/config.env

migrateup:
	migrate -path database/migration -database "postgresql://${user}:${password}@${host}:${port}/${dbname}?sslmode=disable" -verbose up 

migratedown:
	migrate -path database/migration -database "postgresql://${user}:${password}@${host}:${port}/${dbname}?sslmode=disable" -verbose down
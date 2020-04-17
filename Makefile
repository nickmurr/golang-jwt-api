# Server  =========================================
dev:
	cd ./server && go run main.go

build:
	cd ./server && go build .

# Postgres ========================================
pt:
	psql -h localhost -p 5432 -U postgres postgres
pg:
	pgcli -h localhost -p 5432 -U postgres postgres

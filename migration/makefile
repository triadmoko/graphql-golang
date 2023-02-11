table = $(table)
name = $(name)

url=postgres://postgre:secret@127.0.0.1:5431/graphql?sslmode=disable
version=$(version)
migration-up :
	migrate -database "$(url)" -path ./migrations/ up $(version)
	
migration-down :
	migrate -database "$(url)" -path ./migrations/ down $(version)
	
migration-create:
	migrate create -ext sql -dir ./migrations/ -seq $(name)

migration-force:
	migrate -database "$(url)" -path ./migrations/ force $(version)

migration-version:
	migrate -database "$(url)" -path ./migrations/ version
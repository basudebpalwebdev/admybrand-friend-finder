
# AdmybrandFriendFinder assignment

CRUD functionality for a User model and unit test caes along with api doc.




## Documentation

- This Project uses Fiber as web-server framework
- Postgres as DB
- Docker image of Postgres13.3-alpine3.14 used to create the db container
- go-migrate for migrating the tables
- SQLC is used to generate CRUD functionality from SQL queries
- swaggo is used to generate api Doc




  
## Deployment

To deploy this project run

```bash
  docker pull postgres:13.3-alpine3.14 // Download Postgres docker image

  make build-postgres                  // Build the container

  make start-postgres                  // Start the container

  make create-db                       // Create the database

  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  // To install migrate. *add Go home to the path

  go get github.com/kyleconroy/sqlc/cmd/sqlc     // To install SQLC

  go get -u                            // To install required packages

  make migrate-up                      // Run the migrations

  make sqlc                            // Generate CRUD using SQLC

  make test                            // To run test cases

  swag init                            // To generate api docs

  make serve                           // To run the server on port 9999
```

  
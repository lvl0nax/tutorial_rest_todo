### Start Postgres for the app
```
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
```

### Exec psql for the started Postgres container
```
docker exec -it [CONTAINER ID]  psql -U postgres
docker exec -it 0342e999819b psql -U postgres
```

### Migrate DB
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```
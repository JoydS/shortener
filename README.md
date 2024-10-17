### Init
```
docker volume create shortener-postgres-volume
```

### Start
```
docker compose up
```

### Tests
```
go test ./...
```

### Notes :
- Shorten est en POST sur http://localhost:8080/shorten
- Redirect est en GET sur http://localhost:8080/{slug}
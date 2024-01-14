


```sh
docker run --rm -it --name postgres-container -e POSTGRES_USER=username -e POSTGRES_PASSWORD=quote -e POSTGRES_DB=yourdb -p 5432:5432 -d postgres
```

```sh
docker exec -it pgreloserverx_source-db_1 psql -U bookworm -d products
```
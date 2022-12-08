Migrate Db up:
```migrate -path ./storage/migrations -database 'postgres://admin:621@127.0.0.1:5432/category_service?sslmode=disable' up```


Category service da qiladigan ishlarim ro'yxati:

1. postgresda database category_service database ochish  +
2. migration +
    2.1. create table +
    2.2. insert table +
3. storage category +
4. storage product +
5. CRUD category  + 
6. CRUD product + 
7. config setup +
8. .env example + 
9. .env   ++ 
10. Run gRPC in bloomRPC +
11
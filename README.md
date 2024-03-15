## gRPC CRUD

### To run the applications, set the env

#### movies-api

| ENV             |            Desc            |         Sample |
|-----------------|:--------------------------:|---------------:|
| MOVIE_REST_PORT | Port to run the API server |           8080 |
| MOVIE_GRPC_HOST |  Host of the gRPC server   | localhost:9090 |

#### movies-service

| ENV               |            Desc             |    Sample |
|-------------------|:---------------------------:|----------:|
| MOVIE_GRPC_PORT   | Port to run the GRPC server |      8080 |
| MOVIE_DB_HOST     |     Host/ip of database     | localhost |
| MOVIE_DB_PORT     |      Port of database       |      3306 |
| MOVIE_DB_USERNAME |    Username of database     |      root |
| MOVIE_DB_PASSWORD |    Password of database     |  password |
| MOVIE_DB_SCHEMA   |     Schema of database      |    movies |

### To generate proto
```shell
cd common
```
```shell
protoc --go_out=. --go-grpc_out=. movie.proto
```
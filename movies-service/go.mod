module github.com/regmi-bpn/movies-services

go 1.22.1

require (
	github.com/google/uuid v1.6.0
	github.com/regmi-bpn/movie-common v0.0.0
	google.golang.org/grpc v1.62.1
	gorm.io/driver/mysql v1.5.4
	gorm.io/gorm v1.25.7
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/regmi-bpn/movie-common v0.0.0 => ../common

## RPC Database Access Layer by Go

### Description

A rpc database access framework by go with automatic code generation on models, protofiles and services.

Write spec file and run build script, add one line to service.go, then you can run CRUD on databases.

### Requirements & Install

* gRPC for go
* go, version > 1.5
* protobuf
* protobuf go plugin

### Third party

* protobuf
* golang
* gRPC
* gorm
* postgres
* mongodb
* redis

### How to 

1. build a yaml file with name [dbname].yaml under spec/
2. define database in yaml file(see user.yaml sample)
3. run build.sh(will generate files under pb/ proto/ test/ model/ server etc..)
4. register services in service.go

### Start

```
go run seivice.go
```

### Test

```
go run test/user.go
```

### Reference

* [GRPC](http://www.grpc.io/)

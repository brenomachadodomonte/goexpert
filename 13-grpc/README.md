### Generate code from protofile

```shell
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

### Evans

```shell
#install
go install github.com/ktr0731/evans@latest

# connect
evans -r repl

# show packages
show package

# select package
package pb

# select service
service CategoryService

# call
call CreateCategory
```
# QOR example application

This is an example application to show and explain features of QOR.

## Quick Started

```shell
# Get example app
$ go get -u github.com/qor/qor-example

# Setup database
$ mysql -uroot -p
mysql> CREATE DATABASE qor_example;

# Run Application
$ cd $GOPATH/src/github.com/qor/qor-example
$ go run main.go
```

## Generate sample data

```go
$ go run db/seeds/main.go
```

## License

Released under the MIT License.

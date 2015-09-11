# QOR example application

This is an example application to show and explain features of [QOR](http://github.com/qor/qor).

Chat Room: [![Join the chat at https://gitter.im/qor/qor](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/qor/qor?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

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

## DEMO

[demo.getqor.com](http://demo.getqor.com/admin)

## License

Released under the MIT License.

[@QORSDK](https://twitter.com/qorsdk)

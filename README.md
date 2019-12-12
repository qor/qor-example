# QOR Example Application

This is an example application to show and explain features of [QOR](http://getqor.com).

Chat Room: [![Join the chat at https://gitter.im/qor/qor](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/qor/qor?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Quick Start

### Prerequisites

1. If you have not already done so, install a GORM supported database &mdash; mySQL, Postgres or SQLite
2. Install Go version 1.8 or later
3. Configure the GOPATH environment

```bash
        $> export GOPATH=~/go
```

### Get example app

```bash
        $> go get -u github.com/qor/qor-example
```

### Setup the database &mdash; if using the default mysql database:

```bash

        $> mysql -uroot -p
        mysql> CREATE DATABASE qor_example;

```

Or, if using Postgres:

```bash

        $> psql -d postgres
        psql# create role <rolename> with login createdb;
        psql# alter role <rolename> with password '<a strong password>'
        psql Ctrl-d

        $ psql --username=<rolename> -d postgres
        psql# create database <db_name>

        $> cd ${GOPATH}/src/github.com/<account owner folder>qor-example/config/
        $> cp application.example.yml to application.yml
        $> cp database.example.yml database.yml
        $> cp smtp.example.yml smtp.yml

```

Edit database.yml to configure the application to connect to the database

```bash

        $> vim database.yml
        :i edit as below
            db:
                name: <db_name>
                adapter: postgres
                host: localhost
                port: 5432
                user: <rolename>
                password: <a strong password>
        :wq

        $>  
```

__TODO__: Integrate with Hashicorp Vault for management of user credentials

### Run Application

```bash

        $> cd $GOPATH/src/github.com/qor/qor-example
        $> go run main.go # The server listens on port 7000
        $> Ctrl-C # to interupt
```

Alternatively, to start the port on an alternate port

```bash
        $> cd $GOPATH/src/github.com/qor/qor-example
        $> (export PORT=<NNNN> && go run main.go) # Listens on port NNNN
        $> Ctrl-C # to interupt
```

### Generate some sample data

```bash
        $> go get github.com/azumads/faker
        $> cd $GOPATH/src/github.com/qor/qor-example
        $> go run config/db/seeds/main.go config/db/seeds/seeds.go
```

__TODO__: Investigate the warnings about GORM and migrations and the warnings about the FullWidthBannerEditor

### Run tests (Pending)

```bash

        $> go test $(go list ./... | grep -v /vendor/ | grep  -v /db/)
```

## Admin Management Interface

[Qor Example admin configuration](https://github.com/qor/qor-example/blob/master/config/admin/admin.go)

Online Demo Website: [demo.getqor.com/admin](http://demo.getqor.com/admin)

## RESTful API

[Qor Example API configuration](https://github.com/qor/qor-example/blob/master/config/api/api.go)

Online Example APIs:

* Users: [http://demo.getqor.com/api/users.json](http://demo.getqor.com/api/users.json)
* User 1: [http://demo.getqor.com/api/users/1.json](http://demo.getqor.com/api/users/1.json)
* User 1's Orders [http://demo.getqor.com/api/users/1/orders.json](http://demo.getqor.com/api/users/1/orders.json)
* User 1's Order 1 [http://demo.getqor.com/api/users/1/orders/1.json](http://demo.getqor.com/api/users/1/orders/1.json)
* User 1's Orders 1's Items [http://demo.getqor.com/api/users/1/orders.json](http://demo.getqor.com/api/users/1/orders/1/items.json)
* Orders: [http://demo.getqor.com/api/orders.json](http://demo.getqor.com/api/orders.json)
* Products: [http://demo.getqor.com/api/products.json](http://demo.getqor.com/api/products.json)

## License

Released under the MIT License.

[@QORSDK](https://twitter.com/qorsdk)

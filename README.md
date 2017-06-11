# Simple E-Commerce

This application will be built using "Go" for the server-side and "React" for the client-side.

## To Do List
1. Sign Up, Sign In, and Sign Out
2. Setup role for Administrator, Merchant, and User
3. CRUD for product's categories and sub categories
4. ....

## GETTING STARTED - CLIENT

The application runs as an HTTP server at port 3000.

```$xslt
npm start
```

## GETTING STARTED - SERVER

This project uses [Glide](https://glide.sh) as this project's package management.
So, run these following commands:
```$xslt
# install glide (a vendoring and dependency management tool), if you don't have it yet
go get -u github.com/Masterminds/glide

# fetch the dependent packages
cd $GOPATH/bagongkia/e-commerce
make depends   # or "glide up"
```

Create a PostgreSQL database by executing the SQL statements given in the file ```testdata/db.sql```.

Default database connection information:
- database_name   : ecommerce
- username        : postgres
- password        : postgres

For different connection, you may modify file ```models/db.go```

Build and run the application by running these following commands:
```$xslt
go run server.go
```
The application runs as an HTTP server at port 9090: [http://localhost:9090/item-categories](http://localhost:9090/item-categories)
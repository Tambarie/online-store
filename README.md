

# Online-Store API

Online-store-API is a `RESTful-API` that manages a product list for an online store.




## Using Online-store-API

To use this API, clone this repository and follow one of two methods shown below:



### Using Your Local Machine

1. [Install Go](https://golang.org/doc/install)
2. [Installing postgres ](https://www.postgresql.org/download/)
3. Create a database named `online-store`.
4. Install all dependencies using:


```
go mod tidy
```
5. Run the application using:

```
make run
```



## Setting Environmental Variables
An environment variable is a text file containing ``KEY=value`` pairs of your secret keys and other private information. For security purposes, it is ignored using ``.gitignore`` and not committed with the rest of your codebase.

To create, ensure you are in the root directory of the project then on your terminal type:
```
touch .env
```
All the variables used within the project can now be added within the ``.env`` file in the following format:
```
POSTGRES_USER=<user>
POSTGRES_HOST=<localhost>
POSTGRES_PASSWORD=<password>
POSTGRES_DB=<dbname>
POSTGRES_PORT=5432
PORT=9090
```

### Using Docker Compose

1. [Install Docker Compose](https://docs.docker.com/compose/install/)
2. [Pull Postgres Image](https://hub.docker.com/_/postgres)
3. `Enter docker-compose up -d` to build the container


## Tests
Testing is done using the GoMock framework. The ``gomock`` package and the ``mockgen``code generation tool are used for this purpose.
If you installed the dependencies using the command given above, then the packages would have been installed. Otherwise, installation can be done using the following commands:
```
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
```

After installing the packages, run:
```
make mock-product
```

The command above helps to generate mock interfaces from a source file.

To run tests, run:
```
Go the application/test folder. 
```

## Author

* Emmanuel Gbaragbo ([Tambarie](https://github.com/Tambarie)) 🐛






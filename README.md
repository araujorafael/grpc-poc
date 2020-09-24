# How to run
Go to the project root folder and run:

```shell
$ docker-compose up
```

And the products list can be retived trought `GET: localhost:8080/products`)

If needed a pgadmin can be used to manipulate the database (`localhost:3000`) _(user: "admin", password: "admin")_


## Project tree
```
├── databases // database scripts
├── Dockerfiles
├── protos // proto files schema
└── services
    ├── discount
    └── product
```

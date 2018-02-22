# Golang Sneaker API

## Prerequisites

* [Golang 1.10](https://golang.org/dl/)
* gcc
* [github.com/gorilla/mux](https://github.com/gorilla/mux)
* [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

## Getting Started

Please install `Golang` and `gcc` before getting started.

1. Starting in your `$GOPATH`, install the required go packages:

```
go get -u github.com/gorilla/mux
go get -u github.com/mattn/go-sqlite3
```

2. Clone this repository:

```
git clone https://github.com/ktross/golang-sneaker-api.git src/github.com/ktross/golang-sneaker-api
```

## API Endpoints

| URL | Method | URL Params | Data Params |
| --- | ------ | ---------- | ----------- |
| /sneakers | GET |  |  |
| /sneaker/:id | GET | id=[integer] |  |
| /sneaker | POST |  | name=[string] |
| /sneaker/:id | DELETE | id=[integer] |  |
| /sneaker/:id/true-to-size | POST | id=[integer] | size=[integer] |


## Example HTTP Requests

GET /sneakers
```
curl -X GET \
  http://127.0.0.1:8001/sneakers \
  -H 'cache-control: no-cache'
```

GET /sneaker/:id
```
curl -X GET \
  http://127.0.0.1:8001/sneaker/1 \
  -H 'cache-control: no-cache'
```

POST /sneaker
```
curl -X POST \
  http://127.0.0.1:8001/sneaker \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"name": "Example"}'
```

DELETE /sneaker/:id
```
curl -X DELETE \
  http://127.0.0.1:8001/sneaker/1 \
  -H 'cache-control: no-cache'
```

POST /sneaker/:id/true-to-size
```
curl -X POST \
  http://127.0.0.1:8001/sneaker/1/true-to-size \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"sneakerid": "1", "size": "5"}'
```
# wex-coding-challenge
The project is written in go, using the hexagonal architecture.

## Setup Project

### Requirements

1. Clone this repo
```bash
git clone https://github.com/kevinsantana/wex-coding-challenge.git
```

2. To run this project you need [docker](https://docs.docker.com/) and [docker-compose](https://docs.docker.com/compose/) up and running on your machine.

3. Copy [.env.example](.env.example) to `.env` and export them with
```bash
make envvars
```

### Run
1. Build docker image with
```bash
make docker-build
```

2. Run postgres docker database
```bash
make docker-postgres
```

3. Either export [WEX.postman_collection](docs/postman/WEX.postman_collection.json) or make a curl request
```bash
curl --location 'localhost:3060/api/v1/purchase' \
--header 'Content-Type: application/json' \
--data '{
    "description": "new purchase test",
    "amount": 89.9999
}'
```

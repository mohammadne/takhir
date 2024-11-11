# Takhir

> Takhir is a persian name means delay

Here we mimic a vendoring platform which has an ordering system and in this service we want to handle and investigate delayed items in the delivery system to process them separately

## Setup

### Local

```bash
# setup requirements
docker compose -f cluster/docker-compose.yml up -d

go run main.go migrations up # setup the database
go run main.go server # run the server

# teardown application
CTRL+C # to terminate the go application
docker compose -f cluster/docker-compose.yml down
```

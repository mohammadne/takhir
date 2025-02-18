# Zanbil

> Zanbil is a persian name means basket 🧺

Zanbil is a lightweight and efficient Go program designed to demonstrate how to write, test, build, and release Go applications across multiple platforms, including Docker (Podman) and Kubernetes which follows best practices for error handling, internationalization and clean architecture.

Whether you’re building a small project or an enterprise-level application, Zanbil provides valuable insights into structuring and developing Go programs with scalability and maintainability in mind.

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

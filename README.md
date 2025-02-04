# Ory Keto Golang Demo

A small demo which checks permission for the given object, subject and action.

## How to run

```go run main.go```

```bash
curl "http://localhost:8080/check-permission?namespace=Order&object=111&relation=owner&subject=alice"
```
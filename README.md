# Ory Keto Golang Demo

A small demo which requests to the Ory network to check the predefined permission for the given object, subject and action.

## How to run

```go run main.go```

```bash
curl "http://localhost:8080/check-permission?namespace=Order&object=111&relation=owner&subject=alice"
```

# Banknots exchanger

REST API server to search for all options for exchanging the entered amount from the entered list of banknote denominations.

## API Specification
https://rustam2202.github.io/exchanger/

## Run Locally

Config parameters in ```cmd/config.yaml``` and start servers:

```bash
  make run
```
or
```bash
  make build
  make exe
```

Try on `localhost:8080/docs/index.html`

## Running Tests

```bash
  make test
```

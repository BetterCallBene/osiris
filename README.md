# osiris

## compile and run

### build server

```bash
go run cmd/server/main.go
```

### build compiler

```bash
go run cmd/compiler/main.go
```

### build client

```bash
cd cmd/client
GOOS=js GOARCH=wasm go build -o demo.wasm
```
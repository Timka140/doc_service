# Генерируем структуру передачи

## Windows

```text
   Скачать с git архив для windows
   распаковать в нужную папку
   добавить в PATH
```

```bash
Windows
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/transport/protocol/protocol.proto

protoc --go_out=plugins=grpc:. pkg/transport/protocol/protocol.proto
```

## Linux

```text
    Linux
    1. Install the protocol compiler plugins for Go using the following commands:
        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
        $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    
    2. Update your PATH so that the protoc compiler can find the plugins:
        $ export PATH="$PATH:$(go env GOPATH)/bin"

    3. Compile .proto
    protoc --go_opt=paths=source_relative --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/transport/protocol/protocol.proto
```

nano /etc/environment
/root/go/bin

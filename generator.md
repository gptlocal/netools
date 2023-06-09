
## protoc generator

### common

```bash
# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative your_proto_file.proto
$ protoc --go_out=. --go_opt=paths=source_relative common/log/log.proto

$ protoc --go_out=. --go_opt=paths=source_relative common/net/port.proto
$ protoc --go_out=. --go_opt=paths=source_relative common/net/network.proto
$ protoc --go_out=. --go_opt=paths=source_relative common/net/address.proto
$ protoc --go_out=. --go_opt=paths=source_relative common/net/destination.proto

$ protoc --go_out=. --go_opt=paths=source_relative common/serial/typed_message.proto

$ protoc --go_out=. --go_opt=paths=source_relative common/protocol/headers.proto
$ protoc --go_out=. --go_opt=paths=source_relative common/protocol/user.proto
$ protoc --go_out=. --go_opt=paths=source_relative common/protocol/server_spec.proto

$ go mod tidy
```

### transport

```bash
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/global/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/tcp/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/tls/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/reality/config.proto

$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/headers/noop/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/headers/srtp/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/headers/utp/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/headers/wechat/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative transport/internet/headers/wireguard/config.proto
```

### app/proxy

> For Testing

```bash
$ protoc --go_out=. --go_opt=paths=source_relative app/dispatcher/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative app/proxyman/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative app/stats/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative app/policy/config.proto

$ protoc --go_out=. --go_opt=paths=source_relative proxy/freedom/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative proxy/dokodemo/config.proto
$ protoc --go_out=. --go_opt=paths=source_relative proxy/vmess/account.proto
$ protoc --go_out=. --go_opt=paths=source_relative proxy/vmess/outbound/config.proto
```

### core

```bash
$ protoc --go_out=. --go_opt=paths=source_relative core/config.proto
```

### error generator

### common

```bash
$ go generate ./...
```
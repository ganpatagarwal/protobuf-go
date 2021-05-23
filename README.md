# folder structure

```
.
├── README.md
├── client
│   └── client.go
├── go.mod
├── go.sum
├── records
├── server
│   └── server.go
└── status
    ├── status.pb.go
    └── status.proto
```

## status

contains `.proto` file and generated `.pb.go` file

## server

contains code to write records

## client

contains code to read records

## References

https://developers.google.com/protocol-buffers/docs/gotutorial

https://github.com/protocolbuffers/protobuf/tree/master/examples


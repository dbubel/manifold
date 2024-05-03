#### Setup mac
```shell
brew install protobuf
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```


#### Setup linux
```shell
sudo apt update
sudo apt install -y protobuf-compiler
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

```
#### Make sure you have this in your path
https://github.com/protocolbuffers/protobuf/releases
#### Also make sure you have your go install location in your path 
PATH=#PATH:/goinstall/bin

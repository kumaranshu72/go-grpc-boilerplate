<h1>Live Tracking API</h1>
<p>
This repo consists of the code required for live tracking of user on google maps. This project is build in golang and the api's have been exposed using gRPC protocol.
</p>

<h1>Project Setup Instructions</h1>
<ol>
    <li>Install Docker</li>
    <li>docker-compose up --build</li>
</ol>

<h1>Installing proto cli</h1>
<code>brew install protobuf</code>

<h1>Add Path Variables</h1>
<code>export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
</code>

<h1>Instructions to compile protofile</h1>
<code>protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 health.proto</code>

<h1>Instructions to generate Open SSL certificate</h1>
<code>
$ openssl genrsa -out cert/server.key 2048
$ openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
$ openssl req -new -sha256 -key cert/server.key -out cert/server.csr
$ openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650
</code>

<h1>Instructions to run server</h1>
<code>./cmd/server/server -grpc-port=3000 -http-port=8080 -log-level=-1 -log-time-format=2006-01-02T15:04:05.999999999Z07:00</code>

<h1>Pending action items</h1>
- config managemnent using viper
- storing logs to file
- centralized logging with coorelation id support
- context as singleton
- Authentication middleware for http and grpc
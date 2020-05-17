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
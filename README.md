# jsii webassemly bug reproduction

```
git clone git@github.com:Hunter-Thompson/jsii-webasembly-bug.git
cd jsii-webasembly-bug
GOOS=js GOARCH=wasm go build -o main.wasm main.go
cd server
go build 
cd ..
./server/server -listen :8080
```

1. Go on `localhost:8080`
2. Open console & refresh

# jsii webassemly bug reproduction

```
GOOS=js GOARCH=wasm go build -o main.wasm main.go
cd server
go build 
cd ..
./server/ -listen :8080
```

1. Go on `localhost:8080`
2. Open console & refresh

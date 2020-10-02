A simple file server.

```
GO111MODULE=on go get github.com/broady/fileserver
fileserver :9999 .
```

```
docker run -p 9999:9999 -v $PWD:$PWD --rm ghcr.io/broady/fileserver :9999 $PWD
```

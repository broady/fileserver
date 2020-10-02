A simple file server.

go run github.com/broady/fileserver :9999 .

docker run -p 9090:9090 -v $PWD:$PWD --rm ghcr.io/broady/fileserver :9090 $PWD

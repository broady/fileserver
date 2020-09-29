FROM golang:1.15
WORKDIR /fileserver
ADD . .
RUN go build -o /bin/fileserver

FROM gcr.io/distroless/base
COPY --from=0 /bin/fileserver /bin/fileserver
ENTRYPOINT [ "/bin/fileserver" ]
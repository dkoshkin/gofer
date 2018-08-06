FROM dkoshkin/golang-dev:1.10.3-alpine
WORKDIR /go/src/github.com/dkoshkin/gofer
COPY . /go/src/github.com/dkoshkin/gofer
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/gofer main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates openssl
WORKDIR /gofer
COPY --from=0 /go/src/github.com/dkoshkin/gofer/bin/gofer /usr/local/bin
ENTRYPOINT [ "/usr/local/bin/gofer" ]
CMD ["--help"]
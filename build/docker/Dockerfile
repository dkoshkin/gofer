FROM golang:1.10.3-alpine
ARG VERSION
ARG BUILD_DATE
WORKDIR /go/src/github.com/dkoshkin/gofer
COPY . /go/src/github.com/dkoshkin/gofer
RUN echo $VERSION
RUN CGO_ENABLED=0 GOOS=linux go build \
        -ldflags "-X main.version=$VERSION -X 'main.buildDate=$BUILD_DATE'" \
        -o bin/gofer cmd/cli/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates openssl
WORKDIR /gofer
COPY --from=0 /go/src/github.com/dkoshkin/gofer/bin/gofer /usr/local/bin
ENTRYPOINT [ "/usr/local/bin/gofer" ]
CMD ["--help"]
# Usage:
# docker build -t go-time . && \
# docker run -it --rm -v ${HOME}/.go-time.toml:/root/.go-time.toml go-time

FROM amd64/golang:alpine as builder
RUN mkdir /build
WORKDIR /build
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-time go-time.go

FROM amd64/alpine
RUN apk add -U tzdata
COPY --from=builder /build/go-time /app/
ENTRYPOINT [ "/app/go-time" ]
CMD ["/app/go-time"]

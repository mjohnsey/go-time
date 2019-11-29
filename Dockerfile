# Usage:
# docker build -t go-time . && \
# docker run -it --rm -v ${HOME}/.go-time.toml:/root/.go-time.toml go-time

FROM golang:alpine as builder
RUN mkdir /build
WORKDIR /build
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o go-time go-time.go

FROM alpine
RUN apk add -U tzdata
COPY --from=builder /build/go-time /app/
ENTRYPOINT [ "/app/go-time" ]
CMD ["/app/go-time"]

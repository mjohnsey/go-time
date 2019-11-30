# Usage:
# docker build -t go-time . && \
# docker run -it --rm -v ${HOME}/.go-time.toml:/root/.go-time.toml go-time

FROM arm32v7/golang:alpine as builder
RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOARCH=arm GOARM=7 GOOS=linux CGO_ENABLED=0 go build -a -o go-time go-time.go

FROM arm32v7/alpine
RUN apk add -U tzdata
COPY --from=builder /build/go-time /app/
ENTRYPOINT [ "/app/go-time" ]
CMD ["/app/go-time"]

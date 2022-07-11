FROM golang:1.15-alpine as build_base
RUN apk add bash git gcc g++ libc-dev
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
ENV HOST = "0.0.0.0"
ENV ACCESS_JWT = "ACCESS_JWT"
ENV REFRESH_JWT="REFRESH_JWT"
ENV MODE="prod"
ENV PORT = "8080"

COPY . .
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o bin/ ./cmd/main.go

### Put the binary onto base image
FROM plugins/base:linux-amd64
LABEL maintainer="Hank Kuo <asdf024681029@gmail.com>"
EXPOSE 8080
COPY --from=server_builder /bin /bin
COPY ./sqlite3.db /sqlite3.db 

CMD ["/bin/main"]
FROM golang:latest

WORKDIR /app
ENV GOPATH=/

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o build/film-library server/cmd/film-library" --command=./build/film-library
FROM golang:1.17

WORKDIR /usr/src/app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./src/

CMD ["app"]
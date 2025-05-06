FROM golang:latest


WORKDIR /app

COPY . .

RUN go build -o migrations ./migrations

RUN go build -o app ./cmd

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]

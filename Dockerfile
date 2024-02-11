FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o strees_test_go main.go

EXPOSE 8080

COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]

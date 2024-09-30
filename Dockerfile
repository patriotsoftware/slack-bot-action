# syntax=docker/dockerfile:1

FROM golang:1.23


ENV GOPATH=/go
WORKDIR $GOPATH/src/slack-bot-action

# Copy the source code
COPY . .

RUN go build -o /app 

EXPOSE 8080

CMD ["/app"]

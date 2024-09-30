# syntax=docker/dockerfile:1

FROM golang:1.23


ENV GOPATH=/go
WORKDIR $GOPATH/src/slack-bot-action

# Copy the source code
COPY . .

# Build your Go application
RUN go build -o /app 

# Expose the port your application listens on (if applicable)
EXPOSE 8080

# Run your application
CMD ["/app"]
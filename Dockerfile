FROM golang:latest

# Updates

# Working src
WORKDIR /go/src/app

# Copy app
COPY . .

# Run
RUN go build -o main ./src
EXPOSE 8080
CMD ["./main"]
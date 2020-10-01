FROM golang:latest

# Working src
WORKDIR /go/src/app
ADD src src

# Packages src
RUN go get github.com/Masterminds/glide
ADD glide.yaml glide.yaml
ADD glide.lock glide.lock
RUN glide install

# Run main.go
CMD ["go", "run", "src/main.go"]
FROM golang:latest

# ENV
# ENV APP_NAME httpdocker
# ENV PORT 8080

# Working src
WORKDIR /go/src/app
ADD src src

# Packages src
# RUN go get github.com/Masterminds/glide
# ADD glide.yaml glide.yaml
# ADD glide.lock glide.lock
# RUN glide install

# Go packages
# RUN go get github.com/revel/revel

# Run main.go
# CMD ["go", "run", "src/main.go"]


ENTRYPOINT ["go", "run", "src/main.go"]
EXPOSE 8080
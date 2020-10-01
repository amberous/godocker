### Build & Run
```
docker build -t go-docker .
docker run --rm -it -p 8080:8080 go-docker
```

### SSH
```
docker build -t go-docker .
docker run --rm -it -p 8080:8080 -v $(pwd):/go/src/app go-docker bash
```

### Docker Config
https://hasura.io/blog/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c/
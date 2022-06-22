# docker

Build minimal Docker images with Go binaries.

Build:
```
CGO_ENABLED=0 go build
docker build -f Dockerfile_distroless -t test .

```

Run:
```
docker run -it --rm test
```

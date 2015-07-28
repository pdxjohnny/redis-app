Redis App
---

This is a go application which uses redis

This project was built off of [pdxjohnny/go-starter](https://github.com/pdxjohnny/go-starter)

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

```bash
./script/build
```

Running
---

```bash
./build/redis-app_linux-amd64
docker run --rm -ti redis-app
```


- John Andersen


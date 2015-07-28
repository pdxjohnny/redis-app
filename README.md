Docker Go Base
---

This repo is meant to be used as a starting point for go projects

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

Changing The Name
---


```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen


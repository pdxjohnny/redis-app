FROM busybox
ADD ./build/redis-app_linux-amd64 /app
CMD ["/app"]


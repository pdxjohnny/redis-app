FROM busybox
COPY ./build/redis-app_linux-386 /app
CMD ["/app"]

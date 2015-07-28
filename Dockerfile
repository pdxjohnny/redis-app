FROM ubu
COPY ./build/redis-app_linux-amd64 /app
CMD ["/app"]


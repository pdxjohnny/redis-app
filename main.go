package main

import (
    "fmt"
    "strings"
    "net/http"
    "encoding/json"

    "gopkg.in/redis.v3"
)

var client * redis.Client

func redis_init() {
    client = redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}

func redis_get(key string) []byte {
    value, err := client.Get(key).Result()
    if err == redis.Nil {
        return nil
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println(key, value)
    }
    b, err := json.Marshal(value)
    return b
}

func redis_set(key string, value interface{}) bool {
    err := client.Set(key, value, 0).Err()
    if err != nil {
        panic(err)
        return false
    }
    return true
}

func set_handler(w http.ResponseWriter, r *http.Request) {
    path := strings.Split(r.URL.Path, "/")
    key := path[len(path) - 2]
    value := path[len(path) - 1]
    redis_set(key, value)
    fmt.Fprintf(w, "%v\n", string(redis_get(key)))
}

func get_handler(w http.ResponseWriter, r *http.Request) {
    path := strings.Split(r.URL.Path, "/")
    key := path[len(path) - 1]
    fmt.Fprintf(w, "%v\n", string(redis_get(key)))
    return
}

func main() {
    redis_init()
    fmt.Println("Redis connection initialized")
    http.HandleFunc("/", get_handler)
    http.HandleFunc("/set/", set_handler)
    http.ListenAndServe(":5000", nil)
    fmt.Println("Web Server Running")
}

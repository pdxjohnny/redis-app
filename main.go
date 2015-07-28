package main

import (
    "fmt"
    "net/http"

    "gopkg.in/redis.v3"
)

var client *redis.Client

func redis_init() {
    client = redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    // Output: PONG <nil>
}

func redis_test(w http.ResponseWriter) {
    err := client.Set("key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := client.Get("key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, "key", val)

    val2, err := client.Get("key2").Result()
    if err == redis.Nil {
        fmt.Fprintf(w, "key2 does not exists")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Fprintf(w, "key2", val2)
    }
    // Output: key value
    // key2 does not exists
}

func handler(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    redis_test(w)
    fmt.Println("Finished")
}

func main() {
    redis_init()
    fmt.Println("Redis initialized")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":5000", nil)
}

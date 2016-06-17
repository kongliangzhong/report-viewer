package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "time"
)

var dataService DataService

func main() {
    log.Println("starting creditease report viewer ...")

    // init dataservice:
    dataService = &DataServiceMysql{Url: "reporter:reporter@/report"}
    err := dataService.Init()
    if err != nil {
        log.Fatal(err)
    }
    defer dataService.Close()

    startHttpServer := func(addr string) {
        router := NewRouter()
        log.Fatal(http.ListenAndServe(addr, router))
        log.Println("server stoped !!!")
    }

    if len(os.Args) > 1 {
        if os.Args[1] == "--port" {
            serverAddr := ":9000"
            if len(os.Args) > 2 {
                port := os.Args[2]
                serverAddr = ":" + port
            }
            startHttpServer(serverAddr)
        } else if os.Args[1] == "--import-data" {
            log.Fatal("import data feature not implemented yet.")
        } else {
            log.Fatal("unkown command argument:", os.Args[1])
        }

    } else {
        startHttpServer(":9000")
    }
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
    router.PathPrefix("/assets/").Handler(s)
    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}

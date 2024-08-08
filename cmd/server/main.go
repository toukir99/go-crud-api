package main

import (
    "log"
    "go-crud-api/web/router"
    "net/http"
)

func main() {

    // initialize the mux and set up routing
    mux := http.NewServeMux()
    router.SetupRouter(mux)

    // start server
    log.Printf("Server started on :3000")
    log.Fatal(http.ListenAndServe(":3000", mux))
}

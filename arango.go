package main

import (
    "fmt"
    http "net/http"
    driver "github.com/arangodb/go-driver"
    "github.com/arangodb/go-driver/http"
)

func main() {
	conn, err := http.NewConnection(http.ConnectionConfig{
    Endpoints: []string{"http://100.80.250.207:31718"},
})
if err != nil {
    // Handle error
    fmt.Println("No connection")
}
c, err := driver.NewClient(driver.ClientConfig{
    Connection: conn,
})
if err != nil {
    // Handle error
    fmt.Println("No driver")
}
}

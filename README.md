# go-ftapi

FT Content API client in Golang

# Example

Listen to notifications and print out headlines of updated articles

    package main

    import (
        ftapi "github.com/Financial-Times/go-ftapi"
        "flag"
        "fmt"
        "time"
    )

    var API_KEY = flag.String("key", "", "FT API key")

    func main() {
        flag.Parse()
        client := ftapi.Client{*API_KEY}

        ch_notifications := client.Listen(time.Now().Add(-10 * time.Minute), 1 * time.Minute)

        for {
            select {
            case n := <-ch_notifications:
                if n.Type == "http://www.ft.com/thing/ThingChangeType/UPDATE" {
                    article, err := client.GetArticle(n.ApiUrl)
                    if (err != nil) {
                        panic(err)
                    }
                    fmt.Printf("%s: %s\n", article.PublishedDate, article.Title)
                }
            }
        }
    }


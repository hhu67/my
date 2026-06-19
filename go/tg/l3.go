package main

import (
    "fmt"
    "log"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    doc, err := goquery.NewDocument("https://ru.wikipedia.org/api/rest_v1/page/mobile-html/%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0")
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("img").Each(func(_ int, s *goquery.Selection) {
        src, exists := s.Attr("src")
        if exists {
            fmt.Println(src)
        }
    })
}

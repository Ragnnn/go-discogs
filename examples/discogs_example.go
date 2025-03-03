package main

import (
	"fmt"

	"github.com/irlndts/go-discogs"
)

func main() {
	d, err := discogs.NewClient(&discogs.Options{
		UserAgent: "TestDiscogsClient/0.0.1 +http://example.com",
		Currency:  "USD",
		Token:     "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	release, err := d.Search.Search(discogs.SearchRequest{Q: "middle", PerPage: 3})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", release)
}

package restapi

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Rest struct {
	Header *http.Header
}

func (r *Rest) Get(r_url string) string {
	fmt.Println("----In restapi Get----")

	uri, err := url.Parse(r_url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", uri)

	return "true"
}

func (r *Rest) Post(r_url string) string {
	fmt.Println("----In restapi Post----")

	uri, err := url.Parse(r_url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", uri)

	return "true"
}

func (r *Rest) Put(r_url string) string {
	fmt.Println("----In restapi Put----")

	uri, err := url.Parse(r_url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", uri)

	return "true"
}

func (r *Rest) Delete(r_url string) string {
	fmt.Println("----In restapi Delete----")

	uri, err := url.Parse(r_url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", uri)

	return "true"
}

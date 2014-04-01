package restapi

import (
  "fmt"
  "log"
  "net/url"
)

type Rest struct {

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

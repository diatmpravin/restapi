package restapi

import (
  "fmt"
)

type Rest struct {

}

func (r *Rest) Get(url string) string {
  fmt.Println("----In restapi Get----")
  return url
}

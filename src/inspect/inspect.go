package inspect

import (
  "os"
  "fmt"
  "sort"
  "strconv"
  "net/http"
)

var _ = strconv.Atoi
var _ = fmt.Sprintf
var _ = os.Args

type RootMap map[string]interface{}

var Root = make(RootMap)

var next int = 0

func Register(a interface{}) string {
  next++
  k := fmt.Sprintf("$%d", next)
  Root[k] = a
  return k
}

func (m RootMap) SortedKeys() []string {
  z := make([]string, len(m))
  i := 0
  for k, _ := range m {
    z[i] = k
    i++
  }
  sort.Strings(z)
  return z
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("Content-Type", "text/html; charset=utf-8")
  fmt.Fprintf(w, "Hello World!<p>")
  fmt.Fprintf(w, "<ul>")
  for _, x := range Root.SortedKeys() {
    fmt.Fprintf(w, "<li> <b>%#v ==&gt;</b>  %#v", x, Root[x]);
  }
  fmt.Fprintf(w, "</ul>")
}

func init() {
  http.HandleFunc("/i", handler)
}

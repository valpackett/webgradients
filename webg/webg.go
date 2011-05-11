package webg

import (
  "fmt"
  "http"
//  "image"
//  "image/png"
  "strings"
  "strconv"
)

func getstr(r *http.Request, s string, def string) string {
  val := r.FormValue(s)
  if val != "" {
    res, _ = http.URLUnescape(val)
    return res
  }
  return def
}

func getnum(r *http.Request, s string, def int) int {
  res, _ := strconv.Atoi(getstr(r, s, def))
  return res
}

func getcolor(r *http.Request, s string, def string) {
  return strings.Replace(getstr(r, s, def), "#", "", -1)
}

func init() {
  http.HandleFunc("/make", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  width := getnum(r, "width", 1)
  height := getnum(r, "height", 100)
  start := getcolor(r, "start", "eeeeec")
  end := getcolor(r, "end", "d3d7cf")
  direction := getstr(r, "direction", "down")


  fmt.Fprint(w, width)
}

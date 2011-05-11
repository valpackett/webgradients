package webg
//  Web Gradients
//  Copyright 2011 Grigory V. <me@myfreeweb.ru>
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

import (
//  "fmt"
  "http"
  "image"
  "image/png"
  "encoding/hex"
  "strings"
  "strconv"
)

// Getting values from HTTP GET. DRY, isn't it?
func getstr(r *http.Request, s string, def string) string {
  val := r.FormValue(s)
  if val != "" {
    res, _ := http.URLUnescape(val)
    return res
  }
  return def
}

func getnum(r *http.Request, s string, def int) int {
  res, _ := strconv.Atoi(getstr(r, s, strconv.Itoa(def)))
  return res
}

func getcolor(r *http.Request, s string, def string) string {
  return strings.Replace(getstr(r, s, def), "#", "", -1)
}

// The core. Hardcore.
func hex_to_rgb(s string) image.NRGBAColor {
  b, _ := hex.DecodeString(s)
  return image.NRGBAColor{b[0], b[1], b[2], 0xff}
}
func gradient(i *image.NRGBA, s string, e string, dir string) {
  start := hex_to_rgb(s)
  end := hex_to_rgb(e)
  if dir == "left" || dir = "up" {
    start = end
    end = hex_to_rgb(s)
  }
  height := i.Rect.Max.Y
  width := i.Rect.Max.X
  wh := height
  if dir == "left" || dir == "right" {
    wh = width
  }
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      d := y
      if dir == "left" || dir == "right" {
        d = x
      }
      i.Set(x, y, image.NRGBAColor{
        uint8(int(start.R) + int(float32(d) / float32(wh) * float32(int(end.R) - int(start.R)))),
        uint8(int(start.G) + int(float32(d) / float32(wh) * float32(int(end.G) - int(start.G)))),
        uint8(int(start.B) + int(float32(d) / float32(wh) * float32(int(end.B) - int(start.B)))),
        255})
    }
  }
}

// Whoah!
func init() {
  http.HandleFunc("/make", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  width := getnum(r, "width", 1)
  height := getnum(r, "height", 100)
  start := getcolor(r, "start", "eeeeec")
  end := getcolor(r, "end", "d3d7cf")
  direction := getstr(r, "direction", "down")
  image := image.NewNRGBA(width, height)
  gradient(image, start, end, direction)
  w.Header().Set("Content-Type", "image/png")
  _ = png.Encode(w, image)
}

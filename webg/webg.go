package webg
//  Web Gradients
//  Copyright 2011 Grigory V. <floatboth@me.com>
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
	"io"
	"fmt"
	"url"
	"http"
	"bytes"
	"image"
	"image/png"
	"encoding/hex"
	"strings"
	"strconv"
	"appengine"
	"appengine/memcache"
)

// Getting values from HTTP GET. DRY, isn't it?
func getstr(r *http.Request, s, def string) string {
	val := r.FormValue(s)
	if val != "" {
		res, _ := url.QueryUnescape(val)
		return res
	}
	return def
}

func getnum(r *http.Request, s string, def int) int {
	res, _ := strconv.Atoi(getstr(r, s, strconv.Itoa(def)))
	return res
}

func getcolor(r *http.Request, s, def string) string {
	return strings.Replace(getstr(r, s, def), "#", "", -1)
}

// The core. Hardcore.
func hex_to_rgb(s string) image.NRGBAColor {
	b, _ := hex.DecodeString(s)
	return image.NRGBAColor{b[0], b[1], b[2], 0xff}
}
func gradient(i *image.NRGBA, s, e, dir string) {
	var start, end image.NRGBAColor
	if dir == "left" || dir == "up" {
		start = hex_to_rgb(e)
		end   = hex_to_rgb(s)
	} else {
		start = hex_to_rgb(s)
		end   = hex_to_rgb(e)
	}
	height := &i.Rect.Max.Y
	width  := &i.Rect.Max.X
	var wh *int
	var horiz bool
	if dir == "left" || dir == "right" {
		wh = &width
		horiz = true
	} else {
		wh = &height
		horiz = false
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var d *int
			if horiz == true {
				d = &x
			} else {
				d = &y
			}
			i.Set(x, y, image.NRGBAColor{
				uint8(int(start.R) + int(float32(d)/float32(wh)*float32(int(end.R)-int(start.R)))),
				uint8(int(start.G) + int(float32(d)/float32(wh)*float32(int(end.G)-int(start.G)))),
				uint8(int(start.B) + int(float32(d)/float32(wh)*float32(int(end.B)-int(start.B)))),
				255})
		}
	}
}

// Whoah!
func init() {
	http.HandleFunc("/make", handler)
}

func error(w http.ResponseWriter, t string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, t)
}

func handler(w http.ResponseWriter, r *http.Request) {
	width := getnum(r, "width", 1)
	height := getnum(r, "height", 100)
	if width > 4096 || height > 4096 {
		error(w, "Too big!")
		return
	}
	start := getcolor(r, "start", "eeeeec")
	end := getcolor(r, "end", "d3d7cf")
	direction := getstr(r, "direction", "down")
	cachekey := fmt.Sprintf("%sx%s_%s_%s_%s", strconv.Itoa(width), strconv.Itoa(height), start, end, direction)
	c := appengine.NewContext(r)
	w.Header().Set("Content-Type", "image/png")
	if pic, err := memcache.Get(c, cachekey); err == memcache.ErrCacheMiss {
		buf := new(bytes.Buffer)
		image := image.NewNRGBA(width, height)
		gradient(image, start, end, direction)
		png.Encode(buf, image)
		memcache.Add(c, &memcache.Item{
			Key:   cachekey,
			Value: buf.Bytes(),
		})
		io.Copy(w, buf)
	} else if err == nil {
		w.Write(pic.Value)
	} else {
		error(w, "Some weird error")
	}
}

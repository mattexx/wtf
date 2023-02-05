//go:build ignore

package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

var slugs = make(map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {
    // TODO sep handlers
    if len(r.URL.Path) <= 1 {
        // handle root
        buf, err := ioutil.ReadFile("img/darkside.jpeg")
        if err != nil {
            log.Fatal(err)
        }
        w.Header().Set("Content-Type", "image/jpeg")
        w.Write(buf)
    } else {
        // handle slugs
        slug := r.URL.Path[1:]
        ext, ok := slugs[slug]
        if ok {
            buf, err := ioutil.ReadFile("img/"+slug+"."+ext)
            if err != nil {
                log.Fatal(err)
            }
            w.Header().Set("Content-Type", "image/"+ext)
            w.Write(buf)
        } else {
            // 404 redirect to root
            http.Redirect(w, r, "/", http.StatusFound)
        }
    }
}

func main() {
    // cache image files
    files, err := ioutil.ReadDir("./img/")
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
        split := strings.Split(file.Name(), ".")
        if len(split) != 2 {
            continue
        } else {
            slug := split[0]
            // log.Println(slug)
            ext := split[1]
            // log.Println(ext)
            slugs[slug] = ext
        }
    }
    // handlers
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

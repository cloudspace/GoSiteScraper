package main

import (
	"fmt"
        "github.com/moovweb/gokogiri"
        "io/ioutil"
	"net/http"
	"os"
)

func main () {
        url := os.Args[1]
        resp, err := http.Get(url)
        if err != nil {
         fmt.Println("{\"title\":\"\",\"description\":\"\"}")
         return
        }
        defer resp.Body.Close()
        contents, err := ioutil.ReadAll(resp.Body)
        if err != nil {
         fmt.Println("{\"title\":\"\",\"description\":\"\"}")
         return
        }
        document, _ := gokogiri.ParseHtml(contents)
        root := document.Root()
        titles, _ := root.Search(".//title")
        title := ""
        if len(titles) > 0 {
          title = titles[0].Content()
        }
        metas, _ := root.Search(".//meta[@name='description']")
        description := ""
        if len(metas) > 0 {
          description = metas[0].Attribute("content").Content()
        }
        fmt.Println("{\"title\":\""+title+"\",\"description\":\""+description+"\"}")
}

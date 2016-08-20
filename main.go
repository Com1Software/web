//------------------------------------------------------ web
//------------------ (c) 1992-2016 Com1 Software Development
//----------------------------------------------------------
package main

import(
       "fmt"
       "net/http"
       "net"
       "os"
       "io/ioutil"
)
//-------------------------------------------------------------- web
//--------------------------------------------------------- loadPage
func loadPage(title string) (*Page, error) {
   filename := title
   body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
//--------------------------------------------------------- handler
func handler(w http.ResponseWriter, r *http.Request) {
     url:=r.URL.Path[1:]
     p, _ := loadPage(url)
     fmt.Printf("Loaded Page %s.\n",url)
     fmt.Fprintf(w,"%s",p.Body )
}
//--------------------------------------------------------- Page
type Page struct {
    Title string
    Body  []byte
}
//--------------------------------------------------------- MAIN
func main() {
     host, _ := os.Hostname()
     addrs, _ := net.LookupIP(host)
     fmt.Println("use http://localhost:8001/page (On your PC) ")
     fmt.Println("or")
     fmt.Printf("use http://%s:8001/page (On your LAN)\n",addrs[2])
     http.HandleFunc("/", handler)
     http.ListenAndServe(":8001", nil)
}
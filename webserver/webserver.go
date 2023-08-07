package webserver

import (
	"fmt"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// base for webserver
func init() {
	fmt.Println("Starting Web Server...")
}

// StartWebServer starts the webserver
func StartWebServer() {
	//http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("webserver/public")))
	http.ListenAndServe(":8080", nil)

}

// GenerateMarkdown generates markdown for the webserver
// sanitizes unsafe html with bluemonday
func GenerateMarkdown(rw http.ResponseWriter, req *http.Request) {
	unsafe := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	rw.Write(html)
}

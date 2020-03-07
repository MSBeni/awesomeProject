package dynowebportal

import (
	"fmt"
	"net/http"
)

// ListenAndServe starts an HTTP server with a given address and handler
// RunWebPortal starts running the dino web portal on the addr address
func RunWebPortal(addr string) error {
	//ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to
	// use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
	//http.Handle("/foo", fooHandler)
	//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))})
	//log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/", roothandler)
	//err := http.ListenAndServe(addr, nil)
	//return err
	return http.ListenAndServe(addr, nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the dino web portal %s", r.RemoteAddr)

}

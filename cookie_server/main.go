package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// curl --http1.0 -d title="The Art of Community" -d author="Jono Bacon" http://localhost:18888

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Set-Cookie", "VISIT=TRUE; VISITA=TRUE")
	// w.Header().Add("Set-Cookie", "")
	w.Header().Add("WWW-Authenticate", "Basic")
	w.WriteHeader(401)

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	if _, ok := r.Header["Cookie"]; ok {
		// クッキーがあるということは一度来たことがある人
		fmt.Fprintf(w, "<html><body>2 回目以降 </body></html>\n")
		fmt.Fprintf(w, `
		<html>
			<body>
				<form method="POST" >
					<input name="title">
					<input name="author">
					<input type="submit">
				</form>
			</body>
		</html>
		`)
	} else {
		fmt.Fprintf(w, "<html><body> 初訪問 </body></html>\n")
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		// クッキーがあるということは一度来たことがある人
		fmt.Fprintf(w, "<html><body>2 回目以降 </body></html>\n")
	} else {
		fmt.Fprintf(w, "<html><body> 初訪問 </body></html>\n")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

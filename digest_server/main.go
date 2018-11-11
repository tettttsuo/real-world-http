package main

import (
	// import のセクションに以下の 2 行を追加 "io/ioutil"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

func handlerDigest(w http.ResponseWriter, r *http.Request) {
	pp.Printf("URL: %s\n", r.URL.String())
	pp.Printf("Query: %v\n", r.URL.Query())
	pp.Printf("Proto: %s\n", r.Proto)
	pp.Printf("Method: %s\n", r.Method)
	pp.Printf("Header: %v\n", r.Header)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("--body--\n%s\n", string(body))
	if _, ok := r.Header["Authorization"]; !ok {
		w.Header().Add("WWW-Authenticate", `Digest realm="Secret Zone", nonce="TgLc25U2BQA=f510a27804 73e18e6587be702c2e67fe2b04afd", algorithm=MD5, qop="auth"`)
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Fprintf(w, "<html><body>secret page</body></html>\n")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handlerDigest)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

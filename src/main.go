package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Print("监听端口6060")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":6060", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	urls, ok := r.URL.Query()["url"]
	if !ok || len(urls) < 1 {
		http.Error(w,"参数错误",500)
		return
	}

	url := urls[0]
	newReq, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w,"创建请求失败",500)
		return
	}
	Request(w, newReq)
}

func Request(rw http.ResponseWriter, req *http.Request) {
	transport := http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *req

	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		rw.Write([]byte(err.Error()))
		return
	}

	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
	res.Body.Close()
}

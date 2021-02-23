package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/molizz/dingding-go"
)

type ToDingDingResp struct {
	MsgSignature string `json:"msg_signature"`
	TimeStamp    string `json:"timeStamp"`
	Nonce        string `json:"nonce"`
	Encrypt      string `json:"encrypt"`
}

type HttpServer struct {
}

func (h *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.String())

	body, err := dingding.NewDingTalkCrypto(
		"EPxW3vTcLtX83hjvnGFy9E30NkTeVjtp61lCONytZxNZjIZL",
		"iziMCs0M3iU1lK5oeFr7Ac3IIPJUnW4IiGDwYmu8y06",
		"dingmkoxn5exihsog95s").GetEncryptMsg("success")

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func main() {
	http.ListenAndServe(":9620", &HttpServer{})
}

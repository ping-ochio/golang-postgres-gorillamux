package trafic

import (
	"fmt"
	"net/http"
)

func GetIpAddress(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	xforward := r.Header.Get("X-Forwarded-For")
	fmt.Println("IP : ", ip)
	fmt.Println("X-Forwarded-For : ", xforward)
}

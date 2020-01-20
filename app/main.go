package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"

	F "./fib"
	S "./sum"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		res := &response{Message: "Hello World"}

		reqIp := getIP(r)
		res.RequesterIP = reqIp

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			res.EnvVars = append(res.EnvVars, pair[0]+"="+pair[1])
		}
		sort.Strings(res.EnvVars)

		f := F.Fib()
		for i := 1; i <= 90; i++ {
			res.Fib = append(res.Fib, f())
		}

		numbers := []int{1, 2, 3, 4, 5}
		add := S.Sum(numbers)
		res.Sum = add

		// Beautify the JSON output
		out, _ := json.MarshalIndent(res, "", "  ")

		// Normally this would be application/json, but we don't want to prompt downloads
		w.Header().Set("Content-Type", "text/plain")

		io.WriteString(w, string(out))

		fmt.Println("Hello world - from: " + reqIp)
	})
	http.ListenAndServe(":8080", nil)
}

type response struct {
	RequesterIP string   `json:"requesterIP"`
	Message     string   `json:"message"`
	EnvVars     []string `json:"env"`
	Fib         []int    `json:"fib"`
	Sum         int      `json:"sum"`
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

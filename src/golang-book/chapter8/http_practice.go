/* HTTP */

package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head> 
				<title>Hello, World</title>
			</head>
			<body> 
				I am Ibrahim's first HTTP server in Go!!!!
			</body>
		 </html>`,
	)
}

// Copy and paste http://localhost:9000/hello on your browser

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}

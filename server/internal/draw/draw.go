package draw

import "net/http"

func DrawHandler(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "./static/draw.html")
}

package encodetrans

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/axgle/mahonia"
)

//NewResponse wrapped http.ResponseWriter
type NewResponse struct {
	http.ResponseWriter
}

// func (res *NewResponse) WriteHeader(code int) {
// 	res.ResponseWriter.WriteHeader(code)
// }

// func (res *NewResponse) Header() http.Header {
// 	return res.ResponseWriter.Header()
// }

func (res *NewResponse) Write(b []byte) (int, error) {
	dec := mahonia.NewEncoder("GB18030")
	newContent := dec.ConvertString(string(b))
	bytes := []byte(newContent)

	res.Header().Set("Content-Type", http.DetectContentType(bytes))
	return res.ResponseWriter.Write(bytes)
}

func encodingConversion(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		dec := mahonia.NewDecoder("GB18030")
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newContent := dec.ConvertString(buf.String())
		r.Body = ioutil.NopCloser(strings.NewReader(newContent))
		r.ContentLength = int64(len(newContent))

		res := NewResponse{w}
		next.ServeHTTP(&res, r)
	})
}

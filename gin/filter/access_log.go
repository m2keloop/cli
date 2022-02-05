package filter

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type httpReq struct {
	get  url.Values
	data string
}

func AccessLog() gin.HandlerFunc {
	return func(g *gin.Context) {

		cacheBuf := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: g.Writer,
		}
		g.Writer = cacheBuf

		contentType := g.ContentType()
		params := new(httpReq)

		var (
			data []byte
			err  error
		)

		if contentType == "multipart/form-data" {
			g.Request.ParseMultipartForm(32 << 20)
			data, _ = json.Marshal(
				g.Request.PostForm)
		} else if contentType == "application/json" {
			data, err = g.GetRawData()
			if err != nil {
				return
			}
			g.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
		if len(data) > 0 {
			params.data = string(data)
		}

		g.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		//now := time.Now()
		g.Next()
		//log.Info("access log: method:%v, http_code:%v, elapsed:%v", g.Request.RequestURI, cacheBuf.Status(), time.Since(now))

		params.get = g.Request.URL.Query()
		//log4go.WithContext(g).Info("request: %+v", params)
		//log4go.WithContext(g).Info("Response: %v", cacheBuf.body.String())
	}
}

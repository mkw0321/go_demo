package forward_proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Pxy struct{} //自定义HTTP处理器

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport //定义一个数据连接池
	//1.浅拷贝
	outReq := new(http.Request)
	*outReq = *req
	//解析出客户端IP
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP) //与当前clientIP拼接  表示多个代理服务器之间的转发路径
	}
	//2.请求下游
	res, err := transport.RoundTrip(outReq) //将请求发送到目标服务器
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	//3.把下游请求到的内容返回给上游
	for k, value := range outReq.Header {
		for _, v := range value {
			rw.Header().Add(k, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, req.Body)
	res.Body.Close() //关闭连接池
}

func main() {
	fmt.Println("Serve on :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}

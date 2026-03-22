package urlrequest

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/url"
)

func executeRequest(u *url.URL) ([]byte, error) {
	var conn net.Conn
	var err error
	var fullPath string
	path := u.Path

	if len([]byte(path)) == 0 {
		path = "/"
	}

	fullPath = path

	if len([]byte(u.RawQuery)) > 0 {
		fullPath += "?" + u.RawQuery
	}

	if u.Scheme == "http" {
		conn, err = net.Dial("tcp", u.Host+":80")
	} else {
		conn, err = tls.Dial("tcp", u.Host+":443", &tls.Config{})
	}

	if err != nil {
		return nil, fmt.Errorf("Could not open a tcp connection with %s", u.Host)
	}
	defer conn.Close()

	request := fmt.Sprintf("GET %s HTTP/1.1\r\n"+
		"Host: %s\r\n"+
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\r\n"+
		"User-agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Mobile Safari/537.36\r\n"+
		"Connection: close\r\n\r\n", fullPath, u.Host)
	conn.Write([]byte(request))

	rawResponse, err := io.ReadAll(conn)
	return rawResponse, err
}

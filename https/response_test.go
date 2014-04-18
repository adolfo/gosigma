// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package https

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func createTestResponse(code int) (*http.Response, error) {
	msg := fmt.Sprintf(`HTTP/1.1 %d SomeMessage
Server: cloudflare-nginx
Date: Fri, 18 Apr 2014 08:19:45 GMT
Transfer-Encoding: chunked
Connection: keep-alive
Set-Cookie: __cfduid=d1cccd774afa317fcc34389711346edb61397809184442; expires=Mon, 23-Dec-2019 23:50:00 GMT; path=/; domain=.cloudsigma.com; HttpOnly
X-API-Version: Neon.prod.06e860ef2cb5+
CF-RAY: 11cf706acb32088d-FRA

`, code)
	return createResponseFromString(msg)
}

func createTestResponseWithType(code int, contentType string) (*http.Response, error) {
	msg := fmt.Sprintf(`HTTP/1.1 %d SomeMessage
Server: cloudflare-nginx
Date: Fri, 18 Apr 2014 08:19:45 GMT
Content-Type: %s
Transfer-Encoding: chunked
Connection: keep-alive
Set-Cookie: __cfduid=d1cccd774afa317fcc34389711346edb61397809184442; expires=Mon, 23-Dec-2019 23:50:00 GMT; path=/; domain=.cloudsigma.com; HttpOnly
X-API-Version: Neon.prod.06e860ef2cb5+
CF-RAY: 11cf706acb32088d-FRA

`, code, contentType)
	return createResponseFromString(msg)
}

func createResponseFromString(s string) (*http.Response, error) {
	r := bufio.NewReader(strings.NewReader(s))
	return http.ReadResponse(r, nil)
}

func TestHttpsResponseVerifyNoContentType(t *testing.T) {
	hr, err := createTestResponse(200)
	if err != nil {
		t.Error(err)
		return
	}

	r := Response{hr}

	if err := r.VerifyCode(200); err != nil {
		t.Error(err)
	}
	if err := r.VerifyCode(201); err == nil {
		t.Error("expects no error, received:", err)
	}

	if err := r.VerifyContentType(""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.VerifyContentType("application/binary"); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.Verify(200, ""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.Verify(201, ""); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.VerifyJSON(200); err == nil {
		t.Error("expects error, received no error")
	}
	if err := r.VerifyJSON(201); err == nil {
		t.Error("expects error, received no error")
	}
}

func TestHttpsResponseVerifyEmptyContentType(t *testing.T) {
	hr, err := createTestResponseWithType(200, "")
	if err != nil {
		t.Error(err)
		return
	}

	r := Response{hr}

	if err := r.VerifyCode(200); err != nil {
		t.Error(err)
	}
	if err := r.VerifyCode(201); err == nil {
		t.Error("expects no error, received:", err)
	}

	if err := r.VerifyContentType(""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.VerifyContentType("application/binary"); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.Verify(200, ""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.Verify(201, ""); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.VerifyJSON(200); err == nil {
		t.Error("expects error, received no error")
	}
	if err := r.VerifyJSON(201); err == nil {
		t.Error("expects error, received no error")
	}
}

func TestHttpsResponseContentType1(t *testing.T) {
	hr, err := createTestResponseWithType(200, "application/json")
	if err != nil {
		t.Error(err)
		return
	}

	r := Response{hr}

	if err := r.VerifyCode(200); err != nil {
		t.Error(err)
	}
	if err := r.VerifyCode(201); err == nil {
		t.Error("expects no error, received:", err)
	}

	if err := r.VerifyContentType(""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.VerifyContentType("application/binary"); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.Verify(200, ""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.Verify(201, ""); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.VerifyJSON(200); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.VerifyJSON(201); err == nil {
		t.Error("expects error, received no error")
	}
}

func TestHttpsResponseContentType2(t *testing.T) {
	hr, err := createTestResponseWithType(200, "application/json; charset=utf-8")
	if err != nil {
		t.Error(err)
		return
	}

	r := Response{hr}

	if err := r.VerifyCode(200); err != nil {
		t.Error(err)
	}
	if err := r.VerifyCode(201); err == nil {
		t.Error("expects no error, received:", err)
	}

	if err := r.VerifyContentType(""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.VerifyContentType("application/binary"); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.Verify(200, ""); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.Verify(201, ""); err == nil {
		t.Error("expects error, received no error")
	}

	if err := r.VerifyJSON(200); err != nil {
		t.Error("expects no error, received:", err)
	}
	if err := r.VerifyJSON(201); err == nil {
		t.Error("expects error, received no error")
	}
}

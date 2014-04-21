// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package https

import (
	"testing"

	"github.com/Altoros/gosigma/https/httpstest"
)

func TestHttpsResponseVerifyNoContentType(t *testing.T) {
	hr, err := httpstest.CreateResponse(200)
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
	hr, err := httpstest.CreateResponseWithType(200, "")
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
	hr, err := httpstest.CreateResponseWithType(200, "application/json")
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
	hr, err := httpstest.CreateResponseWithType(200, "application/json; charset=utf-8")
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

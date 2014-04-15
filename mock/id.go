// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"errors"
	"net/http"
	"strconv"
)

var chId chan int = make(chan int)

func init() {
	go func() {
		for i := 0; ; i++ {
			chId <- i
		}
	}()
}

func genID() int {
	return <-chId
}

var goSigmaId string = http.CanonicalHeaderKey("gosigma-id")
var errorNotFound error = errors.New("Gosigma-Id not found")

func GetID(h http.Header) (int, error) {
	if v, ok := h[goSigmaId]; ok && len(v) > 0 {
		return strconv.Atoi(v[0])
	}
	return -1, errorNotFound
}

func GetIDFromRequest(r *http.Request) int {
	if id, err := GetID(r.Header); err == nil {
		return id
	}
	return genID()
}

func GetIDFromResponse(r *http.Response) int {
	if id, err := GetID(r.Header); err == nil {
		return id
	}
	return genID()
}

func SetID(h http.Header, id int) {
	h.Add(goSigmaId, strconv.Itoa(id))
}

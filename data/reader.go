// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"errors"
	"io"
)

// ReadJson from io.Reader to the interface
func ReadJson(r io.Reader, v interface{}) error {
	dec := json.NewDecoder(r)
	for {
		err := dec.Decode(v)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

type failReader struct{}

func (failReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package mock

import (
	"net/http"
	"net/http/httptest"
)

type JournalEntry struct {
	Name     string
	Request  *http.Request
	Response *httptest.ResponseRecorder
}

var journal map[int][]JournalEntry = make(map[int][]JournalEntry)

type JournalRequest struct {
	id    int
	entry JournalEntry
	reply chan []JournalEntry
}

var chPut chan JournalRequest = make(chan JournalRequest)
var chGet chan JournalRequest = make(chan JournalRequest)

func init() {
	go func() {
		for {
			select {
			case req := <-chPut:
				journal[req.id] = append(journal[req.id], req.entry)
			case req := <-chGet:
				req.reply <- journal[req.id]
			}
		}
	}()
}

func recordJournal(name string, r *http.Request, rr *httptest.ResponseRecorder) {
	id := GetIDFromRequest(r)
	SetID(rr.HeaderMap, id)
	PutJournal(id, name, r, rr)
}

func PutJournal(id int, name string, r *http.Request, rr *httptest.ResponseRecorder) {
	entry := JournalEntry{name, r, rr}
	jr := JournalRequest{id, entry, nil}
	chPut <- jr
}

func GetJournal(id int) []JournalEntry {
	ch := make(chan []JournalEntry)
	jr := JournalRequest{id, JournalEntry{}, ch}
	chGet <- jr
	return <-ch
}

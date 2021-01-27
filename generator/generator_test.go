package main

import (
	"testing"
)

func TestGetQuestions(t *testing.T) {
	InitEndpoints()
	r := getAllQuestions(MustGet(ALL_QUESTION))
	if r != nil {
		t.Logf("total: %d", len(r.Data.AllQuestions))
	}
}

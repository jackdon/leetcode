package main

import (
	"testing"
)

func TestGetQuestions(t *testing.T) {
	InitEndpoints()
	r := getAllQuestions(MustGet(ALL_QUESTION))
	if r != nil {
		if *lang == "en" {
			t.Logf("total: %d", len(r.Data.AllQuestions))
		} else if *lang == "zh" {
			t.Logf("total: %d", len(r.Data.AllQuestionsBeta))
		}
	}
}

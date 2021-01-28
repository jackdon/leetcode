package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func TestGenEntries(t *testing.T) {
	var dir = "/Users/xulingming/Public/workspace/leetcode-2021/website/docs"
	fis, _ := ioutil.ReadDir(dir)
	var arr []string
	for i := 0; i < len(fis); i++ {
		fis2, _ := ioutil.ReadDir(dir + "/" + fis[i].Name())
		for j := 0; j < len(fis2); j++ {
			arr = append(arr, fmt.Sprintf("%s/%s", fis[i].Name(), fis2[j].Name()))
		}
	}
	ioutil.WriteFile(dir+"/"+"list.txt", []byte(strings.Join(arr, "\n")), os.ModePerm)
}

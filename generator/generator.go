package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TruthHun/html2md"
)

var output = flag.String("output", "/home/bx/Public/learning/leetcode-2021/.output", "")
var input = flag.String("input", "/home/bx/Public/learning/leetcode-2021/problem-set.all.json", "")

type stat struct {
	QuestionID int `json:"question_id"` // : 1882,
	// question__article__live               string `json:"question__article__live"`               // : null,
	// question__article__slug               string `json:"question__article__slug"`               // : null,
	// question__article__has_video_solution string `json:"question__article__has_video_solution"` // : null,
	QuestionTitle      string `json:"question__title"`      // : "The Number of Employees Which Report to Each Employee",
	QuestionTitleSlug  string `json:"question__title_slug"` // : "the-number-of-employees-which-report-to-each-employee",
	QuestionHide       bool   `json:"question__hide"`       // : false,
	TotalAcs           int    `json:"total_acs"`            // : 502,
	TotalSubmitted     int    `json:"total_submitted"`      // : 958,
	FrontendQuestionID int    `json:"frontend_question_id"` // : 1731,
	IsNewQuestion      bool   `json:"is_new_question"`      // : true
}

type difficulty struct {
	level int `json:"level"`
}

type StatStatusPairs struct {
	Stat       stat       `json:"stat"`
	Difficulty difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
}
type ProblemSet struct {
	StatStatusPairs []StatStatusPairs `json:"stat_status_pairs"`
}

type GQLParam struct {
	TitleSlug string
}

type GQLResult struct {
	Data struct {
		Question struct {
			CodeSnippets []map[string]string `json:"codeSnippets"`
			Content      string              `json:"content"`
			Title        string              `json:"title"`     // : "Longest Palindromic Substring",
			TitleSlug    string              `json:"titleSlug"` // : "longest-palindromic-substring",
		} `json:"question"`
	} `json:"data"`
}

func main() {
	log.SetOutput(os.Stdout)
	ps, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	arr := new(ProblemSet)
	if err := json.Unmarshal(ps, arr); err != nil {
		os.Exit(0)
	}
	log.Printf("total problems: %d", len(arr.StatStatusPairs))

	gqlTmpl := template.Must(template.New("gql").Parse(GQL))

	log.Println("start download...")
	for _, p := range arr.StatStatusPairs {
		if p.PaidOnly {
			log.Printf(">> require premium: %s, skip.......", p.Stat.QuestionTitle)
			continue
		}
		var bw bytes.Buffer
		if err := gqlTmpl.Execute(&bw, GQLParam{TitleSlug: p.Stat.QuestionTitleSlug}); err == nil {
			if resp, err := http.DefaultClient.Post("https://leetcode.com/graphql", "application/json", &bw); err == nil {
				log.Println(">> request ok")
				if d, err := ioutil.ReadAll(resp.Body); err == nil {
					log.Println(">> read body ok")
					result := new(GQLResult)
					if err := json.Unmarshal(d, result); err == nil {
						log.Printf(">> parse json ok: %s\n", result.Data.Question.Title)
						dir := filepath.Join(*output, fmt.Sprintf("%d", p.Stat.FrontendQuestionID))
						os.Mkdir(dir, os.ModePerm)
						ioutil.WriteFile(filepath.Join(dir, p.Stat.QuestionTitleSlug+".md"), []byte(html2md.Convert(strings.ReplaceAll(result.Data.Question.Content, "\n", "<br>"))), 0777)
					} else {
						log.Printf("parse json err: %v", err)
					}
				}
				resp.Body.Close()
			} else {
				log.Printf("response err: %v", err)
			}
		} else {
			log.Panicf("%v", err)
			break
		}
		time.Sleep(time.Second)
	}
}

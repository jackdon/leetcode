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

var output = flag.String("output", "/Users/xulingming/Public/workspace/leetcode-2021/.output", "")
var input = flag.String("input", "/Users/xulingming/Public/workspace/leetcode-2021/problem-set.all.json", "")
var outputFmt = flag.Int("output-fmt", 0, "")
var lang = flag.String("lang", "zh", "")

type stat struct {
	QuestionID int `json:"question_id"` // : 1882,
	// question__article__live               string `json:"question__article__live"`               // : null,
	// question__article__slug               string `json:"question__article__slug"`               // : null,
	// question__article__has_video_solution string `json:"question__article__has_video_solution"` // : null,
	QuestionTitle      string      `json:"question__title"`      // : "The Number of Employees Which Report to Each Employee",
	QuestionTitleSlug  string      `json:"question__title_slug"` // : "the-number-of-employees-which-report-to-each-employee",
	QuestionHide       bool        `json:"question__hide"`       // : false,
	TotalAcs           int         `json:"total_acs"`            // : 502,
	TotalSubmitted     int         `json:"total_submitted"`      // : 958,
	FrontendQuestionID interface{} `json:"frontend_question_id"` // : 1731,
	IsNewQuestion      bool        `json:"is_new_question"`      // : true
}

type difficulty struct {
	Level int `json:"level"`
}

type StatStatusPairs struct {
	Stat       stat       `json:"stat"`
	Difficulty difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
}
type ProblemSet struct {
	StatStatusPairs []StatStatusPairs `json:"stat_status_pairs"`
}

type GqlAllQuestion struct {
	Data struct {
		AllQuestionsBeta []struct {
			CategoryTitle      string      `json:"categoryTitle"`      // : "Algorithms"
			Difficulty         string      `json:"difficulty"`         // : "Easy"
			IsPaidOnly         bool        `json:"isPaidOnly"`         // : false
			QuestionFrontendId string      `json:"questionFrontendId"` // : "1"
			QuestionId         string      `json:"questionId"`         // : "1"
			Status             interface{} `json:"status"`             // : null
			Title              string      `json:"title"`              // : "Two Sum"
			TitleSlug          string      `json:"titleSlug"`          // : "two-sum"
			TranslatedTitle    string      `json:"translatedTitle"`    // : null
			Typename           string      `json:"__typename"`         // : "RawQuestionNode"
		} `json:"allQuestionsBeta"`
		AllQuestions []struct {
			CategoryTitle      string      `json:"categoryTitle"`      // : "Algorithms"
			Difficulty         string      `json:"difficulty"`         // : "Easy"
			IsPaidOnly         bool        `json:"isPaidOnly"`         // : false
			QuestionFrontendId string      `json:"questionFrontendId"` // : "1"
			QuestionId         string      `json:"questionId"`         // : "1"
			Status             interface{} `json:"status"`             // : null
			Title              string      `json:"title"`              // : "Two Sum"
			TitleSlug          string      `json:"titleSlug"`          // : "two-sum"
			TranslatedTitle    string      `json:"translatedTitle"`    // : null
			Typename           string      `json:"__typename"`         // : "RawQuestionNode"
		} `json:"allQuestions"`
	} `json:"data"`
}

type GQLParam struct {
	TitleSlug string
}

type GQLResult struct {
	Data struct {
		Question struct {
			CodeSnippets      []map[string]string `json:"codeSnippets"`
			Content           string              `json:"content"`
			Title             string              `json:"title"`     // : "Longest Palindromic Substring",
			TitleSlug         string              `json:"titleSlug"` // : "longest-palindromic-substring",
			TranslatedContent string              `json:"translatedContent"`
			TranslatedTitle   string              `json:"translatedTitle"`
		} `json:"question"`
	} `json:"data"`
}

func main() {
	flag.Parse()
	log.SetOutput(os.Stdout)
	// 初始化endpoints配置
	InitEndpoints()

	arr := getAllProblems(MustGet(ALL_PROBLEM))

	log.Printf("total problems: %d", len(arr.StatStatusPairs))

	// 参数模板
	qe := MustGet(QUESTION_DATA)
	gqlTmpl := template.Must(template.New("gql").Parse(qe.GQLTmpl))

	log.Println("start download...")
	for i, _ := range arr.StatStatusPairs {
		p := &arr.StatStatusPairs[i]
		if p.PaidOnly {
			log.Printf(">> require premium: %s, skip.......", p.Stat.QuestionTitle)
			continue
		}
		r := getQuestionData(qe, p, gqlTmpl)
		if r == nil {
			continue
		}
		writeOuput(p, r)
		time.Sleep(time.Millisecond * 200)
	}
}

func getAllProblems(ep *Endpoint) *ProblemSet {
	resp, err := http.DefaultClient.Get(ep.URL)
	if err == nil {
		defer resp.Body.Close()
		if d, err := ioutil.ReadAll(resp.Body); err == nil {
			log.Println(">> read body ok")
			result := new(ProblemSet)
			if err := json.Unmarshal(d, result); err == nil {
				return result
			} else {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
	return nil
}

func getQuestionData(ep *Endpoint, p *StatStatusPairs, gqlTmpl *template.Template) *GQLResult {
	var bw bytes.Buffer
	if err := gqlTmpl.Execute(&bw, GQLParam{TitleSlug: p.Stat.QuestionTitleSlug}); err == nil {
		if resp, err := http.DefaultClient.Post(ep.URL, "application/json", &bw); err == nil {
			log.Println(">> request ok")
			if d, err := ioutil.ReadAll(resp.Body); err == nil {
				resp.Body.Close()
				log.Println(">> read body ok")
				result := new(GQLResult)
				if err := json.Unmarshal(d, result); err == nil {
					log.Printf(">> parse json ok: %s\n", result.Data.Question.Title)
					return result
				} else {
					log.Printf("parse json err: %v", err)
				}
			}
		} else {
			log.Printf("response err: %v\n", err)
		}
	} else {
		log.Printf("%v\n", err)
	}
	return nil
}

func writeOuput(p *StatStatusPairs, qr *GQLResult) error {
	var dir, file string
	id, slug := p.Stat.FrontendQuestionID, p.Stat.QuestionTitleSlug
	if *outputFmt == 0 {
		dir = filepath.Join(*output, fmt.Sprintf("%v", id))
		file = slug + ".md"
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			if os.IsExist(err) {
				log.Println("dir exists")
			} else {
				log.Printf("create dir unkonwn err: %v", err)
			}
			return err
		}
	} else {
		dir = *output
		file = fmt.Sprintf("%v.%s.md", id, slug)
	}
	{
		var content = ""
		if *lang == "zh" {
			content = qr.Data.Question.TranslatedContent
		} else {
			content = qr.Data.Question.Content
		}
		ioutil.WriteFile(filepath.Join(dir, file), []byte(html2md.Convert(strings.ReplaceAll(content, "\n", "<br>"))), 0777)
		return nil
	}
}

func getAllQuestions(ep *Endpoint) *GqlAllQuestion {
	resp, err := http.DefaultClient.Post(ep.URL, "application/json", bytes.NewReader([]byte(ep.GQLTmpl)))
	if err == nil {
		defer resp.Body.Close()
		if d, err := ioutil.ReadAll(resp.Body); err == nil {
			log.Println(">> read body ok")
			result := new(GqlAllQuestion)
			if err := json.Unmarshal(d, result); err == nil {
				return result
			} else {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
	return nil
}

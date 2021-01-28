package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/TruthHun/html2md"
)

var output = flag.String("output", "/Users/xulingming/Public/workspace/leetcode-2021/.output", "")
var input = flag.String("config", "/Users/xulingming/Public/workspace/leetcode-2021/problem-set.all.json", "")
var outputFmt = flag.Int("output-fmt", 0, "")
var lang = flag.String("lang", "zh", "")
var addTitle = flag.Bool("title", false, "")

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

func (sp *StatStatusPairs) GetCategoryTitle() string {
	return "Algorithms"
}
func (sp *StatStatusPairs) GetIsPaidOnly() bool {
	return sp.PaidOnly
}
func (sp *StatStatusPairs) GetDifficulty() int {
	return sp.Difficulty.Level
}
func (sp *StatStatusPairs) GetQuestionFrontendId() string {
	return fmt.Sprintf("%v", sp.Stat.FrontendQuestionID)
}
func (sp *StatStatusPairs) GetQuestionId() string {
	return fmt.Sprintf("%v", sp.Stat.QuestionID)
}
func (sp *StatStatusPairs) GetTitle() string {
	return sp.Stat.QuestionTitle
}
func (sp *StatStatusPairs) GetTitleSlug() string {
	return sp.Stat.QuestionTitleSlug
}
func (sp *StatStatusPairs) GetTranslatedTitle() string {
	return sp.Stat.QuestionTitle
}

type ProblemSet struct {
	StatStatusPairs []StatStatusPairs `json:"stat_status_pairs"`
}

type QuestionInfoGetter interface {
	GetCategoryTitle() string
	GetIsPaidOnly() bool
	GetDifficulty() int
	GetQuestionFrontendId() string
	GetQuestionId() string
	GetTitle() string
	GetTitleSlug() string
	GetTranslatedTitle() string
}

type QuestionData struct {
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
}

func (qd *QuestionData) GetCategoryTitle() string {
	return qd.CategoryTitle
}
func (qd *QuestionData) GetIsPaidOnly() bool {
	return qd.IsPaidOnly
}
func (qd *QuestionData) GetDifficulty() int {
	switch qd.Difficulty {
	case "Easy":
		return 1
	case "Medium":
		return 2
	case "Hard":
		return 3
	default:
		return -1
	}
}
func (qd *QuestionData) GetQuestionFrontendId() string {
	return qd.QuestionFrontendId
}
func (qd *QuestionData) GetQuestionId() string {
	return qd.QuestionId
}
func (qd *QuestionData) GetTitle() string {
	return qd.Title
}
func (qd *QuestionData) GetTitleSlug() string {
	return qd.TitleSlug
}
func (qd *QuestionData) GetTranslatedTitle() string {
	return qd.TranslatedTitle
}

type GqlAllQuestion struct {
	Data struct {
		AllQuestionsBeta []QuestionData `json:"allQuestionsBeta"`
		AllQuestions     []QuestionData `json:"allQuestions"`
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

	// arr := getAllProblems(MustGet(ALL_PROBLEM))

	// log.Printf("total problems: %d", len(arr.StatStatusPairs))

	// 参数模板
	qe := MustGet(QUESTION_DATA)
	gqlTmpl := template.Must(template.New("gql").Parse(qe.GQLTmpl))

	allQs := getAllQuestions(MustGet(ALL_QUESTION))
	var qs []QuestionData
	if allQs == nil {
		log.Printf("Get questions failed.")
		return
	}
	if *lang == "zh" {
		qs = allQs.Data.AllQuestionsBeta
	} else {
		qs = allQs.Data.AllQuestions
	}
	log.Printf("total problems: %d", len(qs))
	log.Println("start download...")
	manifest, _ := os.Create(filepath.Join(*output, "./mainifest.txt"))
	for i := range qs {
		q := &qs[i]
		if q.GetIsPaidOnly() {
			log.Printf(">> require premium: %s, skip.......", q.GetTitle())
			continue
		}
		r := getQuestionData(qe, q, gqlTmpl)
		if r == nil {
			continue
		}
		writeOuput(q, r, manifest)
		time.Sleep(time.Millisecond * 50)
	}
	defer manifest.Close()
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

func getQuestionData(ep *Endpoint, p QuestionInfoGetter, gqlTmpl *template.Template) *GQLResult {
	var bw bytes.Buffer
	if err := gqlTmpl.Execute(&bw, GQLParam{TitleSlug: p.GetTitleSlug()}); err == nil {
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

var numReg = regexp.MustCompile(`^\d+$`)

func writeOuput(p QuestionInfoGetter, qr *GQLResult, w io.Writer) error {
	var dir, jsonDir, file string
	id, slug := p.GetQuestionFrontendId(), p.GetTitleSlug()
	if category := p.GetCategoryTitle(); category != "" {
		dir = filepath.Join(*output, category)
		// json dir
		jsonDir = filepath.Join(dir, "json")
		if _, err := os.Stat(jsonDir); os.IsNotExist(err) {
			if err := os.MkdirAll(jsonDir, os.ModePerm); err == nil {
				log.Printf("目录: %s 创建成功\n", dir)
			} else {
				log.Printf("create dir unkonwn err: %v", err)
			}
		}
	}
	filePrefix := ""
	isNum := numReg.MatchString(id)
	if isNum {
		idNum, _ := strconv.Atoi(id)
		filePrefix += fmt.Sprintf("%04d", idNum) + "."
	}
	if *lang == "zh" {
		file = filePrefix + p.GetTranslatedTitle()
	} else {
		file = filePrefix + slug
	}
	w.Write([]byte(p.GetCategoryTitle() + "/" + slug + "\n"))
	{
		var content = ""
		if *lang == "zh" {
			content = qr.Data.Question.TranslatedContent
		} else {
			content = qr.Data.Question.Content
		}
		var docTitle = ""
		docTitle = fmt.Sprintf("---\nid: %s\ntitle: %s\n---\n", slug, file)
		ioutil.WriteFile(filepath.Join(dir, file+".md"), []byte(docTitle+html2md.Convert(strings.ReplaceAll(content, "\n", "<br>"))), 0777)
		if d, err := toJsonBytes(&(qr.Data.Question)); err == nil {
			ioutil.WriteFile(filepath.Join(jsonDir, file+".src.json"), d, os.ModePerm)
		}
		if d, err := toJsonBytes(p); err == nil {
			ioutil.WriteFile(filepath.Join(jsonDir, file+".json"), d, os.ModePerm)
		}
		return nil
	}
}

func toJsonBytes(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
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

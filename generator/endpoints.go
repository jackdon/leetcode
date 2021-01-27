package main

import (
	"errors"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

const (
	ALL_PROBLEM   = "ALL_PROBLEM"
	ALL_QUESTION  = "ALL_QUESTION"
	QUESTION_DATA = "QUESTION_DATA"
)

type Endpoint struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	GQLTmpl string `yaml:"gql"`
}

type Endpoints map[string]map[string]*Endpoint

var leetCodeEndpoints Endpoints

func InitEndpoints() {
	if *input != "" {
		if b, err := ioutil.ReadFile(*input); err == nil {
			if err := yaml.Unmarshal(b, &leetCodeEndpoints); err == nil {
				log.Println("config parse success.")
			} else {
				log.Println("config failed to parse.")
			}
		} else {
			log.Println("config failed to read.")
		}
	}
}

func MustGet(key string) *Endpoint {
	ep, _ := leetCodeEndpoints[*lang][key]
	if ep == nil {
		panic(errors.New("no endpoint!"))
	}
	return ep
}

package parser

import (
	"regexp"
	"reptile/engine"
	"reptile/model"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)

var heightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)cm</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re * regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

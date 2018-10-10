package parser

import (
	"HqCrawler/engine"
	"regexp"
	"strconv"
	"HqCrawler/model"
)

const profileRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var ageRe  = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var sexRe  = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var incomeRe  = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var catRe  = regexp.MustCompile(`<td><span class="label">是否购车：</span>([^<]+)</td>`)
var genderRe  = regexp.MustCompile(`<td><span class="label">性别：</span>([^<]+)</td>`)

func ParserProfile(contents []byte,name string) engine.ParserResult  {


	profile := model.Profile{}

	profile.Name = name
	age,err:= strconv.Atoi(extracString(contents,ageRe))
	if err != nil {
		age = 0
	}
	profile.Age = age
	profile.Gender = extracString(contents,sexRe)

	profile.Marriage = extracString(contents,marriageRe)

	profile.Gender = extracString(contents,genderRe)

	profile.Income = extracString(contents,incomeRe)

	profile.Car = extracString(contents,catRe)

	result := engine.ParserResult{
		Items:[]interface{}{profile},
	}


	return result
}
func extracString(contents []byte,regex *regexp.Regexp) string  {

	match := regex.FindSubmatch(contents)

	if len(match) >=2  {
		return string(match[1])

	}
	return ""
}
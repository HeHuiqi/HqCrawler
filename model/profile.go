package model

import "encoding/json"

type Profile struct {
	//Url string
	//Id string
	Name string
	Gender string//性别
	Age int
	Height int
	Weight int
	Income string
	Marriage string
	Education string
	Occupation string //职业
	Hukou string
	Xinzuo string
	House string
	Car string
}

func FromJsonObj(o interface{}) (Profile,error)  {
	var profile Profile
	s,err := json.Marshal(o)
	if err != nil {
		return profile,err
	}
	err = json.Unmarshal(s,&profile)
	return profile,nil
}

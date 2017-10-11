package smartjson

import (
	simplejson "github.com/bitly/go-simplejson"
)

type Json struct {
	object *simplejson.Json
	Err    error
}

func Unmarshal(data []byte) (*Json) {
	j := new(Json)
	j.object, j.Err = simplejson.NewJson(data)
	return j
}

func (j *Json) Get(key string) *simplejson.Json {
	if j.Err != nil {
		return emptyJson()
	}

	return j.object.Get(key)
}

func (j *Json) GetJsons(key string) []*simplejson.Json {
	result := []*simplejson.Json{}

	jsons := j.Get(key)

	nodes := jsons.MustArray()
	nodesCount := len(nodes)

	for i := 0; i < nodesCount; i++ {
		result = append(result, jsons.GetIndex(i))
	}

	if len(result) == 0 {
		result = append(result, emptyJson())
	}
	return result
}

func toJson(key string, value interface{}) (*simplejson.Json) {
	j := emptyJson()
	j.Set(key, value)
	return j
}

func emptyJson() (*simplejson.Json) {
	j, _ := simplejson.NewJson([]byte(`{}`))
	return j
}

func errorJson(err error) *simplejson.Json {
	return toJson("message", err.Error())
}

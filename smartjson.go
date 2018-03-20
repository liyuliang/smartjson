package smartjson

type Json struct {
	object *simpleJson
	Err    error
}

func Unmarshal(data []byte) (*Json) {
	j := new(Json)
	j.object, j.Err = newJson(data)
	return j
}

func (j *Json) Get(key string) *simpleJson {
	if j.Err != nil {
		return emptyJson()
	}

	return j.object.Get(key)
}

func (j *Json) GetJsons(key string) (result []*simpleJson) {

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

func toJson(key string, value interface{}) (*simpleJson) {
	j := emptyJson()
	j.Set(key, value)
	return j
}

func emptyJson() (*simpleJson) {
	j, _ := newJson([]byte(`{}`))
	return j
}

func errorJson(err error) *simpleJson {
	return toJson("message", err.Error())
}

func (j *Json) GetArray() (jsons []*simpleJson) {

	nodes := j.object.MustArray()
	nodesCount := len(nodes)
	if nodesCount != 0 {

		for i := 0; i < nodesCount; i++ {
			node := j.object.GetIndex(i)

			childNodeStr := node.MustString()

			childJson, err := newJson([]byte(childNodeStr))
			if err == nil {
				jsons = append(jsons, childJson)

			} else {
				jsons = append(jsons, node)
			}
		}
	}
	return jsons
}

func (j *Json) GetMap() (data map[string]interface{}) {
	return j.object.MustMap()
}

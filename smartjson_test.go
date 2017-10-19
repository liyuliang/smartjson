package smartjson

import (
	"testing"
)

func Test_toJson(t *testing.T) {

	expectResult := "Brinten"
	json := toJson("name", expectResult)
	result := json.Get("name").MustString()
	if result != expectResult {
		t.Error("string to json failed")
	}
}

func Test_Unmarshal1(t *testing.T) {
	data := `{
    "test": {
        "array": [1, "2", 3],
        "int": 10,
        "float": 5.150,
        "bignum": 9223372036854775807,
        "string": "simplejson",
        "bool": true
    }
}`

	result, err := Unmarshal([]byte(data)).Get("test").Get("int").Int()
	if nil != err {
		t.Error(err)
	}
	if 10 != result {
		t.Error("simple json parse the string failed, expect 10 ,but get ", result)
	}
}

func Test_Unmarshal2(t *testing.T) {
	json := Unmarshal([]byte(jsonStrDemo()))
	jsons := json.GetJsons("hotComments")

	if json.Err == nil {

		val := jsons[0].Get("NotExistProperty").MustString()
		if val != "" {
			t.Error("parse json wrong ,which should be empty string because get not exist property")
		}

		likedCount := jsons[0].Get("likedCount").MustInt()
		expectLinkedCount := 112
		if likedCount != expectLinkedCount {
			t.Error("parse json wrong ,which should be ", expectLinkedCount, ",but get: ", likedCount)
		}

		userId := jsons[0].Get("user").Get("userId").MustInt()
		expectUserId := 60224340
		if userId != expectUserId {
			t.Error("parse json wrong ,which should be ", expectUserId, ",but get: ", userId)

		}
	} else {
		t.Error("parse json failed , please check ',' ")
	}
}

func Test_Unmarshal3(t *testing.T) {
	json := Unmarshal([]byte(jsonStrDemo()))
	json.GetArray()
	jsons := json.GetJsons("hotCommentsNotExist")

	if json.Err == nil {

		val := jsons[0].Get("NotExistProperty").MustString()
		if val != "" {
			t.Error("parse json wrong ,which should be empty string because get not exist property")
		}

		likedCount := jsons[0].Get("likedCount").MustInt()
		expectLinkedCount := 0
		if likedCount != expectLinkedCount {
			t.Error("parse json wrong ,which should be ", expectLinkedCount, ",but get: ", likedCount)
		}

		userId := jsons[0].Get("user").Get("userId").MustInt()
		expectUserId := 0
		if userId != expectUserId {
			t.Error("parse json wrong ,which should be ", expectUserId, ",but get: ", userId)

		}
	} else {
		t.Error("parse json failed , please check ',' ")
	}
}

func Test_Unmarshal4(t *testing.T) {

	jsonStr := `["{\"CreateTime\":\"2017-10-17 00:08:10\",\"Data\":\"{\\\"age\\\":\\\"18\\\",\\\"money\\\":\\\"3933\\\",\\\"name\\\":\\\"keke\\\",\\\"site\\\":\\\"qqqqq\\\",\\\"type\\\":\\\"page\\\",\\\"url\\\":\\\"http://www.qqqqq.org/taotao/10946.html\\\"}\",\"DelayTime\":\"2017-10-17 00:08:11\",\"TTL\":\"2017-10-17 02:08:11\",\"id\":\"139952183956213760\"}","{\"CreateTime\":\"2017-10-17 00:08:08\",\"Data\":\"{\\\"age\\\":\\\"18\\\",\\\"money\\\":\\\"3933\\\",\\\"name\\\":\\\"keke\\\",\\\"site\\\":\\\"qqqqq\\\",\\\"type\\\":\\\"page\\\",\\\"url\\\":\\\"http://www.qqqqq.org/taotao/10946.html\\\"}\",\"DelayTime\":\"2017-10-17 00:08:09\",\"TTL\":\"2017-10-17 02:08:09\",\"id\":\"139952079153139712\"}"]`
	json := Unmarshal([]byte(jsonStr))
	for index, value := range json.GetArray() {

		time := value.Get("CreateTime").MustString()

		if index == 0 && "2017-10-17 00:08:10" != time {
			t.Error("parse json wrong ,which should be ", "2017-10-17 00:08:10", ",but get: ", time)

		}
		if index == 1 && "2017-10-17 00:08:08" != time {
			t.Error("parse json wrong ,which should be ", "2017-10-17 00:08:10", ",but get: ", time)

		}
	}
}

func Test_ToMap(t *testing.T) {
	jsonStr := `["{\"CreateTime\":\"2017-10-17 00:08:10\",\"Data\":\"{\\\"age\\\":\\\"18\\\",\\\"money\\\":\\\"3933\\\",\\\"name\\\":\\\"keke\\\",\\\"site\\\":\\\"qqqqq\\\",\\\"type\\\":\\\"page\\\",\\\"url\\\":\\\"http://www.qqqqq.org/taotao/10946.html\\\"}\",\"DelayTime\":\"2017-10-17 00:08:11\",\"TTL\":\"2017-10-17 02:08:11\",\"id\":\"139952183956213760\"}","{\"CreateTime\":\"2017-10-17 00:08:08\",\"Data\":\"{\\\"age\\\":\\\"18\\\",\\\"money\\\":\\\"3933\\\",\\\"name\\\":\\\"keke\\\",\\\"site\\\":\\\"qqqqq\\\",\\\"type\\\":\\\"page\\\",\\\"url\\\":\\\"http://www.qqqqq.org/taotao/10946.html\\\"}\",\"DelayTime\":\"2017-10-17 00:08:09\",\"TTL\":\"2017-10-17 02:08:09\",\"id\":\"139952079153139712\"}"]`
	json := Unmarshal([]byte(jsonStr))
	for _, value := range json.GetArray() {

		data := value.Get("Data").MustString()

		for k, v := range Unmarshal([]byte(data)).GetMap() {
			println(k,"-",v.(string))
		}
		//println(value.Get("Data").MustString())
	}
}

func jsonStrDemo() string {
	return `{
  "moreHot": true,
  "hotComments": [
    {
      "user": {
        "locationInfo": null,
        "userId": 60224340,
        "nickname": "最爱金子晴",
        "userType": 0,
        "authStatus": 0,
        "expertTags": null
      },
      "beReplied": [
        {
          "user": {
            "locationInfo": null,
            "userId": 92137290,
            "nickname": "情緒零碎v",
            "userType": 0,
            "authStatus": 0,
            "expertTags": null
          },
          "content": "我弟弟说我老土。都什么时代了还听周杰伦。我当时一巴掌就甩了过去。不要问我为什么[撇嘴]",
          "status": 0
        }
      ],
      "liked": false,
      "likedCount": 112,
      "time": 1450900184288,
      "commentId": 56037071,
      "content": "我一晚上在杰伦所有歌里都看到你弟被打"
    }
  ],
  "code": 200,
  "total": 10827,
  "more": true
}
`

}

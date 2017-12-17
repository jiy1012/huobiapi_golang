package comm

import (
	"bytes"
	"compress/gzip"
	//"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
	"huobiapi/conf"
)

func Gzip(data []byte) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	//fmt.Println("gzip size:", len(b.Bytes()))
}

func UnGzip(byte []byte) []byte {
	b := bytes.NewBuffer(byte)
	r, _ := gzip.NewReader(b)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	//fmt.Println("ungzip size:", len(undatas))
	return undatas
}

func JsonDecodeString(String string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(String), &jsonMap)
	return jsonMap
}

func JsonDecodeByte(bytes []byte) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal(bytes, &jsonMap)
	return jsonMap
}
func JsonEncodeMapToByte(stringMap map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(stringMap)
	if err != nil {
		return nil
	}
	return jsonBytes
}

func SendDDRobot(robotStruct DDRobotStruct) {
	jsonParams, err := json.Marshal(robotStruct)
	if err != nil {
		return
	}
	jsonString := string(jsonParams)

	payload := strings.NewReader(jsonString)

	req, _ := http.NewRequest("POST", conf.DDRobotWebHook, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

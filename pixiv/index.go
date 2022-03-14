package pixiv

import (
	"encoding/json"
	"fmt"
	"hexo_editor/utils"
	"io/ioutil"
	"net/http"
)

// GetPixivInfo
// @Description: 根据插画路径获取详情
// @param url
func GetPixivInfo(url string) map[string]interface{} {
	var jsonRes map[string]interface{}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	jsonStr := utils.GetAttr(res.Body, "meta-preload-data", "content")
	err = json.Unmarshal([]byte(jsonStr), &jsonRes)
	if err != nil {
		panic("操作失败!")
	}
	return jsonRes
}

// GetImageByte
// @Description: 根据P站插画路径取图
// @param url
// @return []byte
func GetImageByte(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", "https://www.pixiv.net/")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

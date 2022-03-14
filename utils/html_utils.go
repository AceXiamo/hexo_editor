package utils

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
)

func GetAttr(body io.Reader, id string, attr string) string {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic("body有误!")
	}
	var jsonStr string
	var data map[string]interface{}
	jsonStr, _ = doc.Find("#" + id).Attr(attr)
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic("操作失败!")
	}
	return jsonStr
}

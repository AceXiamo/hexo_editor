package utils

import (
	"os"
	"path"
	"path/filepath"
)

var CoverInit string = "# 这里用来存放封面图\n"

func GenCoverYml()  {
	var rootPath, _ = filepath.Abs(path.Dir(os.Args[0]))
	// 封面图
	cover, _ := os.Create(rootPath + "/cover.yml")
	cover.Close()
	coverFile, _ := os.OpenFile(rootPath + "/cover.yml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	coverFile.Write([]byte(CoverInit))
}
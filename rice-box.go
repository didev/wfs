package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1560302801, 0),

		Content: string("<!DOCTYPE html>\n<head>\n    <title>WFS</title>\n\t<meta charset=\"utf-8\">\n\t<link rel=\"icon\" type=\"image/png\" href=\"/assets/img/wfs.png\">\n</head>\n<body>\n<center><br><br><br><t3>쉽게 서버를 탐색하세요!</t3><br><t4>여러분의 웹어플리케이션과 연결해서 사용자의 편의성을 높히세요.</t4><br>\n<br><br><br><a href=\"{{.RootPath}}\"><img src=\"/assets/img/wfs_click.svg\" width=\"100\" height=\"70\"></a></center>"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "nopath.html",
		FileModTime: time.Unix(1560320092, 0),

		Content: string("<!DOCTYPE html>\n<head>\n    <title>WFS</title>\n    <meta charset=\"utf-8\">\n\t<link rel=\"icon\" type=\"image/png\" href=\"/assets/img/wfs.png\">\n</head>\n<body>\n    {{.URLPath}} 경로가 존재하지 않습니다.\n</body>\n</html>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "wfs.html",
		FileModTime: time.Unix(1560320934, 0),

		Content: string("<!DOCTYPE html>\n<head>\n    <title>WFS</title>\n    <meta charset=\"utf-8\">\n\t<link rel=\"icon\" type=\"image/png\" href=\"/assets/img/wfs.png\">\n</head>\n<body>\n{{if .Parent}}\n    <a href=\"{{.Parent}}\">..</a><br>\n{{end}}\n</body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1560319345, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"
			file3, // "nopath.html"
			file4, // "wfs.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`assets/template`, &embedded.EmbeddedBox{
		Name: `assets/template`,
		Time: time.Unix(1560319345, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html":  file2,
			"nopath.html": file3,
			"wfs.html":    file4,
		},
	})
}

package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1560332601, 0),

		Content: string("<!DOCTYPE html>\n<head>\n    <title>WFS</title>\n\t<meta charset=\"utf-8\">\n\t<link rel=\"stylesheet\" href=\"/assets/bootstrap-4/css/bootstrap.min.css\">\n    <link rel=\"stylesheet\" href=\"/assets/css/wfs.css\">\n\t<link rel=\"icon\" type=\"image/png\" href=\"/assets/img/wfs.png\">\n</head>\n<body class=\"bg-light\">\n\n<div class=\"container p-5\">\n\t<div class=\"text-center\">\n\t\t웹브라우저에서 서버를 탐색하세요!<br>\n\t\tdilink로 웹어플리케이션과 연결해서<br>\n\t\t파일을 쉽게 열어보세요.\n\t</div>\n</div>\n\n<div class=\"container p-5\">\n\t<div class=\"text-center\">\n\t\t<a href=\"{{.RootPath}}\" class=\"btn btn-dark btn-lg\">{{.RootPath}}</a>\n\t</div>\n</div>\n\n<!-- Footer -->\n<footer class=\"page-footer bg-secondary\">\n\t<div class=\"footer-copyright text-center text-light align-middle\">© 2019 Copyright\n\t\t<a href=\"https://lazypic.org\" class=\"text-light\">Lazypic</a> & <a href=\"http://www.digitalidea.co.kr\" class=\"text-light\">Digitalidea</a>\n\t</div>\n</footer>\n</body>\n</html>\n\t"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "nopath.html",
		FileModTime: time.Unix(1560332645, 0),

		Content: string("<!DOCTYPE html>\n<head>\n    <title>WFS</title>\n    <meta charset=\"utf-8\">\n    <link rel=\"stylesheet\" href=\"/assets/bootstrap-4/css/bootstrap.min.css\">\n    <link rel=\"stylesheet\" href=\"/assets/css/wfs.css\">\n\t<link rel=\"icon\" type=\"image/png\" href=\"/assets/img/wfs.png\">\n</head>\n<body class=\"bg-light\">\n\n\n<div class=\"container p-5\">\n    <div class=\"text-center\">\n        {{.URLPath}}<br>\n        경로가 존재하지 않습니다.\n    </div>\n</div>\n<!-- Footer -->\n<footer class=\"page-footer bg-secondary\">\n    <div class=\"footer-copyright text-center text-light align-middle\">© 2019 Copyright\n        <a href=\"https://lazypic.org\" class=\"text-light\">Lazypic</a> & <a href=\"http://www.digitalidea.co.kr\" class=\"text-light\">Digitalidea</a>\n    </div>\n</footer>\n</body>\n</html>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "wfs.html",
		FileModTime: time.Unix(1560332699, 0),

		Content: string("<!DOCTYPE html>\n<head>\n    <title>WFS</title>\n    <meta charset=\"utf-8\">\n    <link rel=\"stylesheet\" href=\"/assets/bootstrap-4/css/bootstrap.min.css\">\n    <link rel=\"stylesheet\" href=\"/assets/css/wfs.css\">\n\t<link rel=\"icon\" type=\"image/png\" href=\"/assets/img/wfs.png\">\n</head>\n<body class=\"bg-light\">\n<div class=\"p-3\">\n    {{if .Parent}}\n    <div class=\"row pl-3\">\n        <a href=\"{{.Parent}}\" class=\"text-dark\">..</a>\n    </div>\n    {{end}}\n\n    {{range .Items}}\n        {{if eq .Typ \"directory\"}}\n        <div class=\"row pl-3\">\n            <a href=\"dilink://{{.Path}}\"><img src=\"/assets/img/{{.Typ}}.png\"></a>&nbsp;\n            <a href=\"{{.Path}}\" class=\"text-dark\">{{.Filename}}</a>\n        </div>\n        {{else}}\n        <div class=\"row pl-3\">\n            <img src=\"/assets/img/{{.Typ}}.png\">&nbsp;<a href=\"dilink://{{.Path}}\" class=\"text-dark\">{{.Filename}}</a>\n        </div>\n        {{end}}\n    {{end}}\n</div>\n\n<!-- Footer -->\n<footer class=\"page-footer bg-secondary\">\n    <div class=\"footer-copyright text-center text-light align-middle\">© 2019 Copyright\n        <a href=\"https://lazypic.org\" class=\"text-light\">Lazypic</a> & <a href=\"http://www.digitalidea.co.kr\" class=\"text-light\">Digitalidea</a>\n    </div>\n</footer>\n</body>\n</html>\n"),
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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/digital-idea/dipath"
)

var (
	flagHTTP     = flag.String("http", "", "service port ex):8081")
	flagRootPath = flag.String("rootpath", "/show", "wfs root path")
)

type item struct {
	Typ      string
	Path     string
	Filename string
}
type recipe struct {
	RootPath string
	URLPath  string
	Parent   string
	Items    []item
	Error    string
	Nukefile string
}

// Index 함수는 wfs "/"의 endpoint 함수입니다.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	rcp := recipe{}
	rcp.RootPath = Home2Abspath(*flagRootPath)
	rcp.URLPath = r.URL.Path

	// teamplate 로딩
	templateBox, err := rice.FindBox("assets/template")
	if err != nil {
		log.Fatal(err)
	}

	if rcp.URLPath == "/" {
		templateString, err := templateBox.String("index.html")
		if err != nil {
			log.Fatal(err)
		}
		index, err := template.New("index").Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
		index.Execute(w, rcp)
		return
	}
	// wfs.html 템플릿 사용하기
	templateString, err := templateBox.String("wfs.html")
	if err != nil {
		log.Fatal(err)
	}
	wfs, err := template.New("wfs").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	if dipath.Exist(r.URL.Path) {
		// 상위경로가 존재한다면, 부모경로를 추가한다.
		if len(strings.Split(rcp.URLPath, "/")) > 2 {
			rcp.Parent = filepath.Dir(rcp.URLPath)
		}
		// 폴더에 존재하는 파일을 불러온다.
		files, err := ioutil.ReadDir(rcp.URLPath + "/")
		if err != nil {
			rcp.Error = err.Error()
			log.Println(err)
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), ".") || strings.HasSuffix(f.Name(), "~") || strings.Contains(f.Name(), "autosave") || strings.HasSuffix(f.Name(), ".lnk") || strings.HasSuffix(f.Name(), ".mel") || strings.HasSuffix(f.Name(), ".tmp") {
				continue
			}
			if f.IsDir() || strings.HasPrefix(f.Mode().String(), "L") {
				// 폴더의 경우
				i := item{}
				i.Typ = "directory"
				i.Path = rcp.URLPath + "/" + f.Name()
				i.Filename = f.Name()
				rcp.Items = append(rcp.Items, i)
			} else {
				switch filepath.Ext(f.Name()) {
				case ".mov", ".mp4", ".avi", ".mkv", ".rv":
					i := item{}
					i.Typ = filepath.Ext(f.Name())[1:]
					i.Path = rcp.URLPath + "/" + f.Name()
					i.Filename = f.Name()
					rcp.Items = append(rcp.Items, i)
				case ".nk", ".nknc", ".ntp", ".mb", ".ma", ".blend", ".hip", ".hipnc":
					i := item{}
					i.Typ = filepath.Ext(f.Name())[1:]
					i.Path = rcp.URLPath + "/" + f.Name()
					i.Filename = f.Name()
					rcp.Items = append(rcp.Items, i)
				case ".exr", ".png", ".jpg", ".dpx", ".tga", ".psd":
					i := item{}
					i.Typ = filepath.Ext(f.Name())[1:]
					i.Path = rcp.URLPath + "/" + f.Name()
					i.Filename = f.Name()
					rcp.Items = append(rcp.Items, i)
				case ".txt", ".py", ".pyc", ".go":
					i := item{}
					i.Typ = filepath.Ext(f.Name())[1:]
					i.Path = rcp.URLPath + "/" + f.Name()
					i.Filename = f.Name()
					rcp.Items = append(rcp.Items, i)
				case ".obj", ".3dl", ".cube":
					i := item{}
					i.Typ = filepath.Ext(f.Name())[1:]
					i.Path = rcp.URLPath + "/" + f.Name()
					i.Filename = f.Name()
					rcp.Items = append(rcp.Items, i)
				case ".gz", ".zip", "bz2", ".ttf", ".pdf":
					i := item{}
					i.Typ = filepath.Ext(f.Name())[1:]
					i.Path = rcp.URLPath + "/" + f.Name()
					i.Filename = f.Name()
					rcp.Items = append(rcp.Items, i)
				default:
					i := item{}
					i.Typ = "file"
					i.Path = rcp.URLPath + "/" + f.Name()
					i.Filename = f.Name()
					rcp.Items = append(rcp.Items, i)
				}
			}
		}
		wfs.Execute(w, rcp)
		return
	}
	// 이미 폴더가 있다면 위쪽 조건에 의해서 파일을 브라우징한다.

	// 합성팀 뉴크파일 생성
	if regexpCompTask.MatchString(rcp.URLPath) {
		nkf, err := nkfilename(rcp.URLPath, "")
		if err != nil {
			rcp.Error = err.Error()
			log.Println(err)
			wfs.Execute(w, rcp)
			return
		}
		initNukefile(rcp.URLPath, nkf)
		rcp.Nukefile = nkf
		templateString, err := templateBox.String("createNuke.html")
		if err != nil {
			log.Fatal(err)
		}
		createNuke, err := template.New("createNuke").Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
		createNuke.Execute(w, rcp)
		return
	}
	// 라이팅팀, 환경팀, 모션그래픽 팀은 프리컴프 파일을 생성한다.
	if regexpLightTask.MatchString(rcp.URLPath) || regexpEnvTask.MatchString(rcp.URLPath) || regexpMgTask.MatchString(rcp.URLPath) {
		precompPath := rcp.URLPath + "/precomp"
		nkf, err := nkfilename(rcp.URLPath, "")
		if err != nil {
			rcp.Error = err.Error()
			log.Println(err)
			wfs.Execute(w, rcp)
			return
		}
		initNukefile(precompPath, nkf)
		templateString, err := templateBox.String("createNuke.html")
		if err != nil {
			log.Fatal(err)
		}
		createNuke, err := template.New("createNuke").Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
		rcp.URLPath = precompPath
		rcp.Nukefile = nkf
		createNuke.Execute(w, rcp)
		return
	}
	// 매트팀 프리컴프파일 메시지
	if regexpMatteTask.MatchString(rcp.URLPath) {
		// 매트팀은 폴더만 생성하길 요청함.
		err := mkdirs(rcp.URLPath)
		if err != nil {
			rcp.Error = err.Error()
			log.Println(err)
			wfs.Execute(w, rcp)
			return
		}
		templateString, err := templateBox.String("createMatte.html")
		if err != nil {
			log.Fatal(err)
		}
		createMatte, err := template.New("createMatte").Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
		createMatte.Execute(w, nil)
		return
	}

	// FX팀 프리컴프파일 메시지
	if regexpFxTask.MatchString(rcp.URLPath) {
		precompPath := rcp.URLPath + "/precomp"
		// 메인 합성파일을 생성한다.
		nkf, err := nkfilename(rcp.URLPath, "master")
		if err != nil {
			rcp.Error = err.Error()
			log.Println(err)
			wfs.Execute(w, rcp)
			return
		}
		initNukefile(precompPath, nkf)
		templateString, err := templateBox.String("createNuke.html")
		if err != nil {
			log.Fatal(err)
		}
		createNuke, err := template.New("createNuke").Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
		rcp.URLPath = precompPath
		rcp.Nukefile = nkf
		createNuke.Execute(w, rcp)
		return
	}
	// 경로가 존재하지 않는 경우
	templateString, err = templateBox.String("nopath.html")
	if err != nil {
		log.Fatal(err)
	}
	nopath, err := template.New("nopath").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	nopath.Execute(w, rcp)
}

func main() {
	flag.Parse()
	if *flagHTTP == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	ip, err := serviceIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Web Server Start : http://%s%s\n", ip, *flagHTTP)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assets)))
	http.HandleFunc("/", Index)
	err = http.ListenAndServe(*flagHTTP, nil)
	if err != nil {
		log.Fatal(err)
	}
}

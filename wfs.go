package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/digital-idea/dipath"
	"github.com/shurcooL/httpfs/html/vfstemplate"
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

func (i *item) SupportIcon() {
	switch i.Typ {
	case ".mov", ".mp4", ".avi", ".mkv", ".rv":
		i.Typ = strings.TrimPrefix(i.Typ, ".")
	case ".nk", ".nknc", ".ntp", ".mb", ".ma", ".blend", ".hip", ".hipnc":
		i.Typ = strings.TrimPrefix(i.Typ, ".")
	case ".exr", ".png", ".jpg", ".dpx", ".tga", ".psd":
		i.Typ = strings.TrimPrefix(i.Typ, ".")
	case ".txt", ".py", ".pyc", ".go":
		i.Typ = strings.TrimPrefix(i.Typ, ".")
	case ".obj", ".3dl", ".cube":
		i.Typ = strings.TrimPrefix(i.Typ, ".")
	case ".gz", ".zip", "bz2", ".ttf", ".pdf":
		i.Typ = strings.TrimPrefix(i.Typ, ".")
	default:
		i.Typ = "file"
	}
}

type recipe struct {
	RootPath string
	URLPath  string
	Parent   string
	Items    []item
	Error    string
	Nukefile string
}

// LoadTemplates 함수는 템플릿을 로딩합니다.
func LoadTemplates() (*template.Template, error) {
	t := template.New("")
	t, err := vfstemplate.ParseGlob(assets, t, "/template/*.html")
	return t, err
}

// Index 함수는 wfs "/"의 endpoint 함수입니다.
func Index(w http.ResponseWriter, r *http.Request) {
	user, err := user.Current()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := LoadTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	rcp := recipe{}
	rcp.RootPath = Home2Abspath(*flagRootPath)
	rcp.URLPath = r.URL.Path

	if rcp.URLPath == "/" {
		err = t.ExecuteTemplate(w, "index.html", rcp)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	if !strings.HasPrefix(rcp.URLPath, *flagRootPath) {
		err = t.ExecuteTemplate(w, "nopath.html", rcp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
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
				i := item{}
				i.Typ = filepath.Ext(f.Name())
				i.SupportIcon()
				i.Path = rcp.URLPath + "/" + f.Name()
				i.Path = rcp.URLPath + "/" + f.Name()
				// 경로의 시작이 $HOME 과 같다면 ~문자로 치환한다.
				if strings.HasPrefix(i.Path, user.HomeDir) {
					i.Path = strings.Replace(i.Path, user.HomeDir, "~", 1)
				}
				i.Filename = f.Name()
				rcp.Items = append(rcp.Items, i)
			}
		}
		err = t.ExecuteTemplate(w, "wfs.html", rcp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	// 이미 폴더가 있다면 위쪽 조건에 의해서 파일을 브라우징한다.

	// 합성팀 뉴크파일 생성
	if regexpCompTask.MatchString(rcp.URLPath) {
		nkf, err := nkfilename(rcp.URLPath, "")
		if err != nil {
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "wfs.html", rcp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		initNukefile(rcp.URLPath, nkf)
		rcp.Nukefile = nkf

		err = t.ExecuteTemplate(w, "createNuke.html", rcp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	// 라이팅팀, 환경팀, 모션그래픽 팀은 프리컴프 파일을 생성한다.
	if regexpLightTask.MatchString(rcp.URLPath) || regexpEnvTask.MatchString(rcp.URLPath) || regexpMgTask.MatchString(rcp.URLPath) {
		precompPath := rcp.URLPath + "/precomp"
		nkf, err := nkfilename(rcp.URLPath, "")
		if err != nil {
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "wfs", rcp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		initNukefile(precompPath, nkf)
		rcp.URLPath = precompPath
		rcp.Nukefile = nkf
		err = t.ExecuteTemplate(w, "createNuke.html", rcp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	// 매트팀 프리컴프파일 메시지
	if regexpMatteTask.MatchString(rcp.URLPath) {
		// 매트팀은 폴더만 생성하길 요청함.
		err := mkdirs(rcp.URLPath)
		if err != nil {
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "wfs.html", rcp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		err = t.ExecuteTemplate(w, "createMatte.html", rcp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// 경로가 존재하지 않는 경우
	err = t.ExecuteTemplate(w, "nopath.html", rcp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

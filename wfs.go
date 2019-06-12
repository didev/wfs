package main

import (
	"flag"
	"fmt"
	"io"
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
	flagHTTP     = flag.String("http", "", "service port ex):8080")
	flagRootPath = flag.String("rootpath", "/show", "wfs root path")
)

// Index 함수는 wfs "/"의 endpoint 함수입니다.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	type recipe struct {
		RootPath string
		URLPath  string
		Parent   string
	}
	rcp := recipe{}
	rcp.RootPath = *flagRootPath
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
	io.WriteString(w, headHTML)
	if dipath.Exist(r.URL.Path) {
		// 상위경로가 존재한다면, 부모경로를 추가한다.
		if len(strings.Split(rcp.URLPath, "/")) > 2 {
			rcp.Parent = filepath.Dir(rcp.URLPath)
		}
		// 폴더에 존재하는 파일을 불러온다.
		files, err := ioutil.ReadDir(rcp.URLPath + "/")
		if err != nil {
			log.Println(err)
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), ".") || strings.HasSuffix(f.Name(), "~") || strings.Contains(f.Name(), "autosave") || strings.HasSuffix(f.Name(), ".lnk") || strings.HasSuffix(f.Name(), ".mel") || strings.HasSuffix(f.Name(), ".tmp") {
				continue
			}
			if f.IsDir() || strings.HasPrefix(f.Mode().String(), "L") {
				// 폴더의 경우
				io.WriteString(w, fmt.Sprintf(`<div><a href="dilink://%s"><img src="/assets/img/folder.png"></a> <a href="%s">%s</a> </div>`, rcp.URLPath+"/"+f.Name(), rcp.URLPath+"/"+f.Name(), f.Name()))
			} else {
				switch filepath.Ext(f.Name()) {
				case ".mov", ".mp4", ".avi", ".mkv":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/mov.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".rv":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/rv.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".nk":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/nk.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".mb", ".ma":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/maya.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".exr", ".png", ".jpg", ".dpx", ".tga":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/image.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".psd":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/psd.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".txt":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/text.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".py", ".pyc":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/python.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".go":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/go.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".blend":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/blender.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".obj":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/obj.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".ntp":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/natron.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".3dl", ".cube":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/lut.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".gz", ".zip", "bz2":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/zip.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".hip", ".hipnc":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/houdini.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".ttf":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/ttf.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				case ".pdf":
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/pdf.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				default:
					io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/file.png"> <a href="dilink://%s">%s</a></div>`, rcp.URLPath+"/"+f.Name(), f.Name()))
				}
			}
		}
		return
	}
	if regexpCompTask.MatchString(r.URL.Path) {
		nkf, err := nkfilename(r.URL.Path, "")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		initNukefile(r.URL.Path, nkf)
		io.WriteString(w, "Create new nuke file.")
		io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/nk.png"> <a href="dilink://%s">%s</a></div>`, r.URL.Path+"/"+nkf, nkf))
		return
	}

	if regexpLightTask.MatchString(r.URL.Path) || regexpEnvTask.MatchString(r.URL.Path) || regexpMgTask.MatchString(r.URL.Path) {
		precompPath := r.URL.Path + "/precomp"
		nkf, err := nkfilename(r.URL.Path, "")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		initNukefile(precompPath, nkf)
		io.WriteString(w, "Create new nuke file.")
		io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/nk.png"> <a href="dilink://%s">%s</a></div>`, precompPath+"/"+nkf, nkf))
		return
	}

	if regexpMatteTask.MatchString(r.URL.Path) {
		err := mkdirs(r.URL.Path)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		io.WriteString(w, "pub폴더가 생성되었습니다. F5를 눌러주세요.")
		io.WriteString(w, "precomp컴프 필요시 pub/precomp/SS_0010_matte_v01.nk 파일형태로 수동 생성해주세요.")
		return
	}

	if regexpFxTask.MatchString(r.URL.Path) {
		precompPath := r.URL.Path + "/precomp"
		// 메인 합성파일을 생성한다.
		nkf, err := nkfilename(r.URL.Path, "master")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		initNukefile(precompPath, nkf)
		io.WriteString(w, "Create new nuke file.")
		io.WriteString(w, fmt.Sprintf(`<div><img src="/assets/img/nk.png"> <a href="dilink://%s">%s</a></div>`, precompPath+"/"+nkf, nkf))
		return
	}
	// 경로가 존재하지 않는 경우
	templateString, err := templateBox.String("nopath.html")
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

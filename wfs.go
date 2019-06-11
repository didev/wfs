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

	"github.com/digital-idea/dipath"
)

var (
	flagHTTP     = flag.String("http", "", "service port ex):8080")
	flagRootPath = flag.String("rootpath", "/show", "wfs root path")
	// Templatepath 상수는 템플릿 아이콘이 있는 endpoint 입니다.
	Templatepath = "http://10.0.98.20:8080/template/icon/"
)

// Index 함수는 wfs "/"의 endpoint 함수입니다.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		// index.html 템플릿 사용하기
		io.WriteString(w, rootHTML)
		io.WriteString(w, "<center><br><br><br><t3>쉽게 서버를 탐색하세요!</t3><br><t4>여러분의 웹어플리케이션과 연결해서 사용자의 편의성을 높히세요.</t4><br>")
		io.WriteString(w, `<br><br><br><a href="`+*flagRootPath+`"><img src="http://10.0.98.20:8083/media/wfs_click.svg" width="100" height="70"></a></center>`)
		return
	}
	// wfs.html 템플릿 사용하기
	io.WriteString(w, headHTML)
	if dipath.Exist(r.URL.Path) {
		// 상위경로가 존재한다면, .. 를 프린트하는 코드.
		if len(strings.Split(r.URL.Path, "/")) > 2 {
			io.WriteString(w, fmt.Sprintf(`<a href="%s">..</a><br>`, filepath.Dir(r.URL.Path)))
		}

		files, _ := ioutil.ReadDir(r.URL.Path + "/")
		for _, f := range files {
			if strings.HasPrefix(f.Name(), ".") || strings.HasSuffix(f.Name(), "~") || strings.Contains(f.Name(), "autosave") || strings.HasSuffix(f.Name(), ".lnk") || strings.HasSuffix(f.Name(), ".mel") || strings.HasSuffix(f.Name(), ".tmp") {
				continue
			} else {
				if f.IsDir() || strings.HasPrefix(f.Mode().String(), "L") {
					// 폴더의 경우
					io.WriteString(w, fmt.Sprintf(`<div><a href="dilink://%s"><img src="%s/folder.png"></a> <a href="%s">%s</a> </div>`, r.URL.Path+"/"+f.Name(), Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
				} else {
					switch filepath.Ext(f.Name()) {
					case ".mov", ".mp4", ".avi", ".mkv":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/mov.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".rv":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/rv.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".nk":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".mb", ".ma":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/maya.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".exr", ".png", ".jpg", ".dpx", ".tga":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/image.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".psd":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/psd.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".txt":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/text.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".py", ".pyc":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/python.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".go":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/go.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".blend":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/blender.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".obj":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/obj.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".ntp":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/natron.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".3dl", ".cube":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/lut.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".gz", ".zip", "bz2":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/zip.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".hip", ".hipnc":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/houdini.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".ttf":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/ttf.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					case ".pdf":
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/pdf.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					default:
						io.WriteString(w, fmt.Sprintf(`<div><img src="%s/file.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
					}
				}
			}
		}
	} else {
		if regexpCompTask.MatchString(r.URL.Path) {
			nkf, err := nkfilename(r.URL.Path, "")
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			initNukefile(r.URL.Path, nkf)
			io.WriteString(w, "Create new nuke file.")
			io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+nkf, nkf))
		} else if regexpLightTask.MatchString(r.URL.Path) || regexpEnvTask.MatchString(r.URL.Path) || regexpMgTask.MatchString(r.URL.Path) {
			precompPath := r.URL.Path + "/precomp"
			nkf, err := nkfilename(r.URL.Path, "")
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			initNukefile(precompPath, nkf)
			io.WriteString(w, "Create new nuke file.")
			io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, precompPath+"/"+nkf, nkf))
		} else if regexpMatteTask.MatchString(r.URL.Path) {
			err := mkdirs(r.URL.Path)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			io.WriteString(w, "pub폴더가 생성되었습니다. F5를 눌러주세요.")
			io.WriteString(w, "precomp컴프 필요시 pub/precomp/SS_0010_matte_v01.nk 파일형태로 수동 생성해주세요.")

		} else if regexpFxTask.MatchString(r.URL.Path) {
			precompPath := r.URL.Path + "/precomp"
			// 메인 합성파일을 생성한다.
			nkf, err := nkfilename(r.URL.Path, "master")
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			initNukefile(precompPath, nkf)
			io.WriteString(w, "Create new nuke file.")
			io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, precompPath+"/"+nkf, nkf))
		} else {
			io.WriteString(w, "경로가 존재하지 않습니다. : "+r.URL.Path)
		}
	}
}

func main() {
	flag.Parse()
	if *flagHTTP == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(*flagHTTP, nil)
	if err != nil {
		log.Fatal(err)
	}
}

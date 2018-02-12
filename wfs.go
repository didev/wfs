package main

import (
	"di/dipath"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const Templatepath = "http://10.0.98.20:8080/template/icon/"

func www_root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		io.WriteString(w, rootHTML)
		io.WriteString(w, "<center><br><br><br><t3>쉽게 서버를 탐색하세요!</t3><br><t4>여러분의 웹어플리케이션과 연결해서 사용자의 편의성을 높히세요.</t4><br>")
		io.WriteString(w, `<br><br><br><a href="/show"><img src="http://10.0.98.20:8083/media/wfs_click.svg" width="100" height="70"></a></center>`)
	} else {
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
			} else if regexpLightTask.MatchString(r.URL.Path) || regexpEnvTask.MatchString(r.URL.Path) {
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

			} else if regexpFxTask.MatchString(r.URL.Path) {
				precompPath := r.URL.Path + "/precomp"
				io.WriteString(w, "Create new nuke file.")
				// 메인 합성파일을 생성한다.
				nkf, err := nkfilename(r.URL.Path, "")
				if err != nil {
					io.WriteString(w, err.Error())
					return
				}
				initNukefile(precompPath, nkf)
				io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, precompPath+"/"+nkf, nkf))
				// element별 합성파일을 생성한다.
				q := r.URL.Query()
				elements := strings.Split(q.Get("elements"), ",")
				for _, e := range elements {
					if e == "" {
						// 실제로 query에 elements를 선언하지 않더라도..
						// elements값에는 빈문자열 1개가 들어온다.
						continue
					}
					nkf, err := nkfilename(r.URL.Path, e)
					if err != nil {
						io.WriteString(w, err.Error())
						return
					}
					initNukefile(precompPath, nkf)
					io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, precompPath+"/"+nkf, nkf))
				}
			} else {
				io.WriteString(w, "경로가 존재하지 않습니다. : "+r.URL.Path)
			}
		}
	}
}

func main() {
	portPtr := flag.String("http", "", "service port ex):8080")
	flag.Parse()
	if *portPtr == "" {
		fmt.Println("WFS is WebFileSystem for Digitalidea")
		flag.PrintDefaults()
		os.Exit(1)
	}
	http.HandleFunc("/", www_root)
	http.ListenAndServe(*portPtr, nil)
}

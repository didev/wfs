package main

import (
	"fmt"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"net/http"
	"strings"
	"path/filepath"
	)

const Templatepath = "http://10.0.90.193:8080/template/icon/"

func isPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func parent_path(path string) string {
	rpath := ""
	strlist := strings.Split(path, "/")
	for num, i := range strlist {
		if num == len(strlist) - 1 {
			break
		} else {
			rpath += i + "/"
		}
	}
	return rpath[:len(rpath)-1]
}


func www_root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html",)

	if r.URL.Path == "/" {
		io.WriteString(w, rootHTML)
		//main page
		io.WriteString(w, "<center><br><br><br><t3>쉽게 서버를 탐색하세요!</t3><br><t4>여러분의 웹어플리케이션과 연결해서 사용자의 편의성을 높히세요.</t4><br>")
		io.WriteString(w, `<br><br><br><a href="/show"><img src="http://10.0.90.193:8083/media/wfs_click.svg" width="100" height="70"></a></center>`)
	} else {
		io.WriteString(w, headHTML)
		if isPath(r.URL.Path) {
			// 상위경로가 존재한다면, .. 를 프린트하는 코드.
			if len(strings.Split(r.URL.Path, "/")) > 2  {
				io.WriteString(w, fmt.Sprintf(`<a href="%s">..</a><br>`, parent_path(r.URL.Path) ))
			}

			files, _ := ioutil.ReadDir(r.URL.Path + "/")
			for _, f := range files {
				if strings.HasPrefix(f.Name(), ".") || strings.HasSuffix(f.Name(), "~") || strings.Contains(f.Name(), "autosave") || strings.HasSuffix(f.Name(), ".lnk") || strings.HasSuffix(f.Name(), ".mel") || strings.HasSuffix(f.Name(), ".tmp") {
					continue
				} else {
					if f.IsDir() || strings.HasPrefix(f.Mode().String(), "L"){
						io.WriteString(w, fmt.Sprintf(`<div><a href="dilink://%s"><img src="%s/folder.png"></a> <a href="%s">%s</a> </div>`, r.URL.Path+"/"+f.Name(), Templatepath, r.URL.Path+"/"+f.Name(), f.Name() ))
					} else {
						switch filepath.Ext(f.Name()) {
							case ".mov", ".mp4", ".avi", ".mkv" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/mov.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".rv" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/rv.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".nk" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".mb",".ma" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/maya.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".exr",".png",".jpg",".dpx",".tga" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/image.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".psd" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/psd.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".txt" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/text.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".py",".pyc" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/python.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".go" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/go.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".blend" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/blender.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".obj" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/obj.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".ntp" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/natron.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".3dl",".cube" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/lut.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".gz",".zip","bz2" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/zip.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".hip",".hipnc" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/houdini.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".ttf" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/ttf.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							case ".pdf" : io.WriteString(w, fmt.Sprintf(`<div><img src="%s/pdf.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
							default: io.WriteString(w, fmt.Sprintf(`<div><img src="%s/file.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+f.Name(), f.Name()))
						}
					}
				}
			}
		} else {
			if strings.HasPrefix(r.URL.Path, "/show") && strings.HasSuffix(r.URL.Path, "/comp/dev") && !strings.Contains(r.URL.Path,"assets") && ispath(r.URL.Path) == false {
				io.WriteString(w, "Create new nuke file.")
				//make folder.
				os.MkdirAll(r.URL.Path+"/wip", 0777)
				os.Mkdir(r.URL.Path+"/src", 0777)
				os.Mkdir(r.URL.Path+"/tmp", 0777)
				//make nukefile.
				exec.Command("touch", r.URL.Path+"/"+gennk(r.URL.Path) ).Run()
				io.WriteString(w, fmt.Sprintf(`<div><img src="%s/nk.png"> <a href="dilink://%s">%s</a></div>`, Templatepath, r.URL.Path+"/"+gennk(r.URL.Path), gennk(r.URL.Path)))
			} else {
				io.WriteString(w, "not exist path : " + r.URL.Path)
			}
		}
	}
}

func ispath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func gennk(path string) string {
	pathlist := strings.Split(path, "/")
	return pathlist[5] + "_" + pathlist[6] + "_v01.nk"
}

func main() {
	portPtr := flag.String("http", "", "service port ex):8080")
	flag.Parse()
	if *portPtr == "" {
		fmt.Println("WFS is web file system for digitalidea")
		fmt.Println("Copyright (C) 2015 kimhanwoong")

		flag.PrintDefaults()
		os.Exit(1)
	}
	http.HandleFunc("/", www_root)
	http.ListenAndServe(*portPtr, nil)
}

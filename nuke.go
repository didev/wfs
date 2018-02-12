package main

import (
	"di/dipath"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var regexpCompTask = regexp.MustCompile(`/show/\S+/seq/\S+/\S+/comp/dev$`)
var regexpFxTask = regexp.MustCompile(`/show/\S+/seq/\S+/\S+/fx/dev$`)
var regexpLightTask = regexp.MustCompile(`/show/\S+/seq/\S+/\S+/light/dev$`)
var regexpMatteTask = regexp.MustCompile(`/show/\S+/seq/\S+/\S+/matte/pub$`)
var regexpEnvTask = regexp.MustCompile(`/show/\S+/seq/\S+/\S+/env/dev$`)

// nkfilename 함수는 경로, 앨레멘트 이름으로 뉴크파일명을 생성한다.
func nkfilename(path string, element string) (string, error) {
	seq, err := dipath.Seq(path)
	if err != nil {
		return "", err
	}
	shot, err := dipath.Shot(path)
	if err != nil {
		return "", err
	}
	task, err := dipath.Task(path)
	if err != nil {
		return "", err
	}
	if element != "" {
		task += "_" + element
	}
	return fmt.Sprintf("%s_%s_%s_v01.nk", seq, shot, task), nil
}

// mkdirs는 폴더를 생성하고 task이름까지 idea 권한으로 변경한다.
func mkdirs(path string) error {
	err := os.MkdirAll(path, 0775)
	if err != nil {
		return err
	}
	// Task하위의 특정 위치부터 Task폴더까지 권한이 idea:idea 775 형태가 되어야 한다.
	// 예를 들면 아래 폴더 전부 최초 생성시 idea:idea 775 권한을 가져야 한다.
	// /show/TEMP/seq/SS/SS_0010/fx/dev/precomp
	// /show/TEMP/seq/SS/SS_0010/fx/dev
	// /show/TEMP/seq/SS/SS_0010/fx
	current := path
	for {
		_, err := dipath.Task(current)
		if err != nil {
			break
		}
		dipath.Ideapath(current)
		current = filepath.Dir(current)
	}
	return nil
}

// initNukefile함수는 경로, 뉴크파일명으로 필요한 폴더, 파일을 생성한다.
func initNukefile(path, nkfilename string) error {
	err := mkdirs(path)
	if err != nil {
		return err
	}
	// 파일생성
	f, err := os.Create(path + "/" + nkfilename)
	if err != nil {
		return err
	}
	defer f.Close()
	dipath.Ideapath(path + "/" + nkfilename) // 권한설정 idea:idea 775
	return nil
}

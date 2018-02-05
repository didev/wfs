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

func nkfilename(path string) (string, error) {
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
	return fmt.Sprintf("%s_%s_%s_v01.nk", seq, shot, task), nil
}

func initNukefile(path string) error {
	// 폴더생성
	err := os.MkdirAll(path, 0775)
	if err != nil {
		return err
	}
	// Task하위의 특정 위치부터 Task폴더까지 권한이 idea:idea 775 형태가 되어야 한다.
	current := path
	for i := 1; i <= 10; i++ {
		_, err := dipath.Task(current)
		if err != nil {
			break
		}
		dipath.Ideapath(current)
		current = filepath.Dir(current)
	}
	// 파일생성
	nkf, err := nkfilename(path)
	if err != nil {
		return err
	}
	f, err := os.Create(path + "/" + nkf)
	if err != nil {
		return err
	}
	defer f.Close()
	dipath.Ideapath(nkf) // 권한설정 idea:idea 775
	return nil
}

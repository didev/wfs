package main

import (
	"di/dipath"
	"fmt"
	"os"
	"regexp"
)

var regexpCompTask = regexp.MustCompile(`/show/\S+/seq/\S+/\S+/comp/dev$`)
var regexpFxTask = regexp.MustCompile(`/show/\S+/seq/\S+/\S+/fx/dev$`)

func nkfilename(filepath string) (string, error) {
	seq, err := dipath.Seq(filepath)
	if err != nil {
		return "", err
	}
	shot, err := dipath.Shot(filepath)
	if err != nil {
		return "", err
	}
	task, err := dipath.Task(filepath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s_%s_%s_v01.nk", seq, shot, task), nil
}

func initNukefile(filepath string) error {
	// 폴더생성
	err := os.MkdirAll(filepath, 0775)
	if err != nil {
		return err
	}
	dipath.Ideapath(filepath) // 권한설정 idea:idea 775
	// 파일생성
	nkf, err := nkfilename(filepath)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath + "/" + nkf)
	if err != nil {
		return err
	}
	defer f.Close()
	dipath.Ideapath(nkf) // 권한설정 idea:idea 775
	return nil
}

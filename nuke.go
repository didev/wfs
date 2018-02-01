package main

import (
	"di/dipath"
	"fmt"
	"os"
	"regexp"
)

var regexpCompTask = regexp.MustCompile(`/show[/_]/\S+/seq/\S+/\S+/comp/dev$`)
var regexpOtherTask = regexp.MustCompile(`/show[/_]/\S+/seq/\S+/\S+/\S+/dev/precomp$`)

func nkfilename(filepath string) (string, error) {
	shot, err := dipath.Shot(filepath)
	if err != nil {
		return "", err
	}
	task, err := dipath.Task(filepath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s_%s_v01.nk", shot, task), nil
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

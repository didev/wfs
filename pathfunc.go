package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/digital-idea/dipath"
)

var (
	regexpCompTask  = regexp.MustCompile(*flagRootPath + `/\S+/seq/\S+/\S+/comp/dev$`)
	regexpLightTask = regexp.MustCompile(*flagRootPath + `/\S+/seq/\S+/\S+/light/dev$`)
	regexpMatteTask = regexp.MustCompile(*flagRootPath + `/\S+/seq/\S+/\S+/matte/pub$`)
	regexpEnvTask   = regexp.MustCompile(*flagRootPath + `/\S+/seq/\S+/\S+/env/dev$`)
	regexpMgTask    = regexp.MustCompile(*flagRootPath + `/\S+/seq/\S+/\S+/mg/dev$`)
)

// Home2Abspath 함수는 ~ 문자로 경로가 시작하면 물리적인 경로로 바꾸어준다.
func Home2Abspath(p string) (string, error) {
	if !strings.HasPrefix(p, "~") {
		return p, nil
	}
	usr, err := user.Current()
	if err != nil {
		return p, err
	}
	return usr.HomeDir + strings.TrimPrefix(p, "~"), nil
}

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
	err := os.MkdirAll(path, 0775) // 입력받은 경로를 생성한다.
	if err != nil {
		return err
	}
	// Task하위의 특정 위치부터 Task폴더까지 권한이 idea:idea 775 형태가 되어야 한다.
	// 예를 들면 아래 폴더 전부 최초 생성시 idea:idea 775 권한을 가져야 한다.
	// "/show/TEMP/seq/SS/SS_0010/comp/dev/source" 만약 이러한 경로가 들어왔다면 경로에 775를 준다.
	// "/show/TEMP/seq/SS/SS_0010/comp/dev" source를 제거하고 775를 준다.
	// "/show/TEMP/seq/SS/SS_0010/comp" dev를 제거하고 775를 준다.
	// "/show/TEMP/seq/SS/SS_0010" task를 구할 수 없는 경로가 되었으니 권한 설정을 종료한다.
	currentPath := path
	for {
		_, err := dipath.Task(currentPath) // Task를 가지고 올 수 없다면 연산을 멈춘다.
		if err != nil {
			break
		}
		dipath.Ideapath(currentPath)
		currentPath = filepath.Dir(currentPath) // 뒤의 경로를 하나 제거하고 다시 권한을 줄 준비를 한다.
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
	// 권한설정 idea:idea 775
	err = dipath.Ideapath(path + "/" + nkfilename)
	if err != nil {
		return err
	}
	return nil
}

# WFS(Web File System)

![travisCI](https://secure.travis-ci.org/digital-idea/wfs.png)

웹을 통해서 서버의 파일을 브라우징합니다.
필요시 작업에 필요한 뉴크파일을 초기 셋팅합니다.

### 다운로드
- Linux: https://github.com/digital-idea/wfs/releases/download/v1.0/wfs_linux_x86-64.tgz
- MacOS: https://github.com/digital-idea/wfs/releases/download/v1.0/wfs_darwin_x86-64.tgz
- Windows: https://github.com/digital-idea/wfs/releases/download/v1.0/wfs_windows_x86-64.tgz

### 서버실행

```bash
$ wfs -http :8081 -rootpath /show
```

### HISTORY
- '19.6.7 : 오픈소스 전환
- '15.6.10 : 1차 개발완료. 목적 : 웹에서 스토리지의 파일 프리뷰 및 [dilink](https://github.com/digital-idea/dilink) 연동.

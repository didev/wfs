# WFS(Web File System)
- 웹프로토콜을 이용해서 빠른 스피드로 원하는 파일까지 접근
- 또한 클라이언트에서 프로그램을 실행하기 위한 웹 응용프로그램 입니다.
- 합성팀 추가기능으로 comp/dev를 만나면 폴더가 생성되기 때문에 아래처럼 idea로 서비스를 실행합니다.
```
alias wfsd='nohup su - idea -c "/lustre/INHouse/CentOS/bin/wfs -http=:8081 &"'
```

#### HISTORY
- '15.8.10 ~ '15.8.17 : CSI와 함께 특허등록을 위한 문서작성.
- '15.6.10 : 1차 개발완료. 목적 : 웹에서 스토리지의 파일 프리뷰 및 dilink 연동.

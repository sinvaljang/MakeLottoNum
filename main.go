package main

/*
negroni : 웹 서버의 라이프사이클을 관리하고 모든 웹 요청을 받아서 처리하는 역할을 한다.
render  : 웹 요청 처리 결과를 다양한 형태로 렌더링한다.
httprouter  :  웹 요청의 URL을 해석해서 적절한 핸들러로 연결해준다.
*/
import (
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

const (
	//애플리케이션에서 사용할 세션의 키 정보
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"
)

var renderer *render.Render

func init() {
	renderer = render.New()
}

func main() {
	lotto := new(lottoNum)

	//라우터 생성
	//사용자로 부터 요청을 처리하기위한
	router := httprouter.New()

	handlerManager(router)

	makeJson(lotto)

	// negroni 미들웨어 생성
	n := negroni.Classic()
	// negroni에 router를 핸들러로 등록
	n.UseHandler(router)

	//웹서버 실행
	n.Run(":3000")
}

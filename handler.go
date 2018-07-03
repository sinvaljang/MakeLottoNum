package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//핸들러정의
//해당 패턴에 대한 핸들러를 등록한다
func handlerManager(router *httprouter.Router) {

	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

		// 렌더러를 사용하여 템플릿 렌더링
		// 맵 map[keytype]valuetype{initial}
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "Simple Chat!"})
	})

	router.GET("/lotto", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "lotto", nil)
	})
}

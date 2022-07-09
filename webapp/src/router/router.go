package router

import "github.com/gorilla/mux"

// Gerar retornar  um router com todas a rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
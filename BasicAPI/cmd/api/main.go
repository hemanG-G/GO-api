package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hemanG-G/GO_MINI_PROJECTS/tree/main/BasicAPI/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service ... ")

	fmt.Println(`

	GGGG   OOO   AAAAA  PPPPP  III
	GG  GG OO  OO AA   AA PP  PP III
   GG      OO  OO AAAAAAA PPPPP  III
   GG   GG OO  OO AA   AA PP     III
	GGGGGG  OOO   AA   AA PP     III  ASCII ART

	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}

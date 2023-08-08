package fiskil

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	Port  string
	Count Sequence
}

func (s *Server) Bootstrap() {
	s.configurePort()
	fmt.Printf("\nstart listening on %v", s.Port)
	http.ListenAndServe(":"+s.Port, s.defaultRoute())
}

func (s *Server) defaultRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		n := s.Count.Next()
		response := fmt.Sprintf(`{ "seq": %d }`, n) + "\n"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))

		//log
		fmt.Print(response)
	})
	return r
}

func (s *Server) configurePort() {
	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = "8080"
	}
	pi, err := strconv.Atoi(p)
	if err != nil || (pi < 1 || pi > 65535) {
		fmt.Println("invalid port, exiting")
		os.Exit(-1)
	}
	s.Port = p
	fmt.Printf("\nset port to %v", s.Port)
}

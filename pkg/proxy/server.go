package proxy

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	conf   Config
	router *mux.Router
	rules  []Rule
	srv    *http.Server
}

func NewServer(conf Config) (*Server, error) {
	rules, err := parseRules(conf.Rules)
	if err != nil {
		return nil, err
	}

	if len(rules) == 0 {
		return nil, errors.New("no redirect rules")
	}

	router := mux.NewRouter()
	return &Server{
		conf:   conf,
		router: router,
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.Port),
			Handler: router,
		},
		rules: rules,
	}, nil
}

func (s *Server) Start() error {
	s.route()
	log.Printf("proxy listen on: [http://localhost:%v]", s.conf.Port)

	return s.srv.ListenAndServe()
}

func (s *Server) route() {
	s.router.PathPrefix("/").HandlerFunc(s.defaultHandler)
}

func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	reqPath := r.URL.Path
	log.Printf("request path: %v", reqPath)

	for _, l := range s.rules {
		if l.re != nil && l.re.MatchString(reqPath) {
			targetURL := l.target + reqPath
			log.Printf("matched pattern: %v, target: %v\n", l.re, targetURL)
			http.Redirect(w, r, targetURL, http.StatusFound)
			return
		}
	}

	http.Redirect(w, r, "https://www.douban.com", http.StatusFound)
}

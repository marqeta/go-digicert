package digicerttest

import (
	"net/http"
	"net/http/httptest"
)

type Server struct {
	Mux         *http.ServeMux
	HTTPServer  *httptest.Server
	DataStore   map[string]interface{}
	BaseURLPath string
}

func NewTestServer(base_url_path string) *Server {
	if base_url_path == "" {
		base_url_path = "/services/v2"
	}

	mux := http.NewServeMux()
	apiHandler := http.NewServeMux()
	apiHandler.Handle(base_url_path+"/", http.StripPrefix(base_url_path, mux))

	return &Server{
		Mux:         mux,
		HTTPServer:  httptest.NewServer(apiHandler),
		BaseURLPath: base_url_path,
	}
}

func (srv *Server) Close() {
	srv.HTTPServer.Close()
}

func (srv *Server) APIUrl() string {
	return srv.HTTPServer.URL + srv.BaseURLPath
}

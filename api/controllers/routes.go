package controllers

import "github.com/vanjoechua/orderbackend/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(s.GetOrders)).Methods("GET")
}



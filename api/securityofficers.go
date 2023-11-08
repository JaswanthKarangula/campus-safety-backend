package api

func (server *Server) SetUpSecurityOfficerRouter() {

	server.router.POST("/admin/createAdmin", server.createNewAdmin)
}

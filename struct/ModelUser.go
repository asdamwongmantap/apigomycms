package modeluser

type (

	// for response list user
	ShowUser struct {
		UserID       int
		UserName     string
		UserFullname string
		UserEmail    string
		UserPassword string
	}
	// for param request user
	ReqUser struct {
		UserID       int
		UserName     string
		UserFullname string
		UserEmail    string
		UserPassword string
	}
)

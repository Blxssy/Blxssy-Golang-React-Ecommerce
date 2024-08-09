package config

const (
	API         = "/api"
	APIUsers    = API + "/users"
	APIUsersID  = APIUsers + "/:id"
	APIAUTH     = API + "/auth"
	APIREGISTER = APIAUTH + "/register"
	APILOGIN    = APIAUTH + "/login"
	APIREFRESH  = APIAUTH + "/refresh"
)

const (
	APIProducts   = API + "/products"
	APIProductsID = APIProducts + "/:id"
)

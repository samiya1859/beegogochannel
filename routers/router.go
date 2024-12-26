package routers

import (
	"catproject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.CatAPIController{})
	beego.Router("/vote", &controllers.CatAPIController{}, "post:Vote")
	beego.Router("/breed/", &controllers.BreedController{}, "get:GetAllBreeds")
	beego.Router("/breed/:breed_id", &controllers.BreedController{}, "get:GetBreedDetails")
	beego.Router("/breed/images/:breed_id", &controllers.BreedController{}, "get:GetBreedImages")
	beego.Router("/fav", &controllers.CatAPIController{}, "post:Favourite")
	beego.Router("/fav/getall", &controllers.CatAPIController{}, "get:GetFavorites")
}

package routers

import (
	"net/http"

	"app/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

var tokenAuth *jwtauth.JWTAuth

func Router() {
	app := chi.NewRouter()

	tokenAuth = jwtauth.New("HS256", []byte("token"), nil)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	app.Use(cors.Handler)

	app.Post("/login", controller.Login)
	app.Post("/signin", controller.SignIn)

	app.Route("/get", func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/card", controller.GetCard)
		r.Get("/main", controller.GetMain)
		r.Get("/cpu", controller.GetCpu)
		r.Get("/ram", controller.GetRam)
		r.Get("/rom", controller.GetRom)
		r.Get("/psup", controller.GetPSup)
		r.Get("/heat_dissipation", controller.GetHeatDissipation)
		r.Get("/build", controller.GetBuild)
	})

	app.Route("/user_post", func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/build", controller.AddBuild)
		r.Post("/delete", controller.DeleteBuild)

	})

	app.Route("/post", func(r chi.Router) {

		r.Post("/card", controller.AddCard)
		r.Post("/main", controller.AddMain)
		r.Post("/cpu", controller.AddCpu)
		r.Post("/ram", controller.AddRam)
		r.Post("/rom", controller.AddRom)
		r.Post("/psup", controller.AddPsup)
		r.Post("/heat_dissipation", controller.AddHeatDissipation)
	})

	app.Route("/update", func(r chi.Router) {

		r.Post("/card", controller.UpdateCard)
		r.Post("/main", controller.UpdateMain)
		r.Post("/cpu", controller.UpdateCpu)
		r.Post("/ram", controller.UpdateRam)
		r.Post("/rom", controller.UpdateRom)
		r.Post("/psup", controller.UpdatePsup)
		r.Post("/heat_dissipation", controller.UpdateHeatDissipation)
	})

	app.Route("/delete", func(r chi.Router) {

		r.Post("/card", controller.DeleteCard)
		r.Post("/main", controller.DeleteMain)
		r.Post("/cpu", controller.DeleteCpu)
		r.Post("/ram", controller.DeleteRam)
		r.Post("/rom", controller.DeleteRom)
		r.Post("/psup", controller.DeletePsup)
		r.Post("/heat_dissipation", controller.DeleteHeatDissipation)
	})

	app.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":3001", app)
}

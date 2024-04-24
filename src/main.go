package main

import (
	"compel/config"
	"compel/controllers"
	"compel/models"
	"encoding/json"
	"log"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func main() {
	loadEnv()
	models.ConnectToDB()
	models.Migrate()
	serveApplication()
}

func loadEnv() {
	err := config.Read()
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())

	}
}

func serveApplication() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(ginI18n.Localize(
		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				lng := context.GetHeader("Accept-Language")
				if lng == "" {
					return defaultLng
				}
				return lng
			},
		),

		ginI18n.WithBundle(&ginI18n.BundleCfg{
			RootPath:         "./helpers/i18n",
			AcceptLanguage:   []language.Tag{language.Spanish, language.English},
			DefaultLanguage:  language.Spanish,
			UnmarshalFunc:    json.Unmarshal,
			FormatBundleFile: "json",
		}),
	))

	controllers.RegisterController(router)
	router.Run(":" + config.CNF.App.Port)
}

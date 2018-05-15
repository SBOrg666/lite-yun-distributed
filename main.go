package main

import (
	"github.com/gin-gonic/gin"
	"github.com/SBOrg666/lite-yun-distributed/utils"
	"github.com/jasonlvhit/gocron"
	"time"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"github.com/tidwall/gjson"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	store := sessions.NewCookieStore([]byte(uuid.Must(uuid.NewV4()).String()))
	store.Options(sessions.Options{MaxAge: 0, HttpOnly: false})
	router.Use(sessions.Sessions("session", store))

	b, err := ioutil.ReadFile("servers.json")
	if err != nil {
		log.Println(err)
	}
	if (len(string(b)) == 0) {
		utils.ServersString = `{"Servers":[]}`
	} else {
		utils.ServersString = string(b)
	}

	result := gjson.Get(utils.ServersString, "Servers")
	utils.ServersMap = make(map[string]gjson.Result)
	for _, val := range result.Array() {
		utils.ServersMap[val.Get("Token").Str] = val
	}

	log.Println(utils.ServersMap)

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.LoadHTMLFiles("./template/index.html",
		"./template/login.html",
		"./template/processes.html",
		"./template/path.html",
		"./template/about.html",
		"./template/authors.html",
		"./template/list.html",
	)

	router.GET("/", utils.CheckLoginIn(), utils.IndexHandler_get)

	LoginGroup := router.Group("/")
	{
		LoginGroup.GET("/login", utils.LoginHandler_get)
		LoginGroup.POST("/login", utils.LoginHandler_post)
	}

	router.GET("/processes.html", utils.CheckLoginIn(), utils.ProcessHandler_get)

	router.GET("/path", utils.CheckLoginIn(), utils.PathHandler_get)

	router.POST("/listServer", utils.CheckLoginIn(), utils.ListServerHandler_post)
	router.POST("/addServer", utils.CheckLoginIn(), utils.AddServerHandler_post)
	router.POST("/deleteServer", utils.CheckLoginIn(), utils.DeleteServerHandler_post)

	AboutGroup := router.Group("/")
	{
		AboutGroup.GET("/about", utils.CheckLoginIn(), utils.AboutHandler_get)
		AboutGroup.GET("/authors", utils.CheckLoginIn(), utils.AuthorsHandler_get)
	}

	router.GET("/list", utils.CheckLoginIn(), utils.ListHandler_get)

	router.POST("/changeToken", utils.CheckLoginIn(), utils.ChangeTokenHandler_post)

	utils.Upload_data = make([]uint64, 5)
	utils.Download_data = make([]uint64, 5)
	utils.InitUpload = 0
	utils.InitDownload = 0
	utils.Current_Month = int(time.Now().Month())
	gocron.Every(1).Day().Do(utils.UpdateNetworkData)

	router.Run(":8000")
}

package utils

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"log"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"strings"
)

type User struct {
	Name     string `form:"username"`
	Password string `form:"password"`
}

type FileList struct {
	Files []string `json:"files"`
}

type Server struct {
	Token    string
	Ip       string
	Port     string
	Username string
	Password string
}

type PathInfo struct {
	Path     string
	Dirs     []DirItem
	Files    []FileItem
	Writable bool
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}

func logErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandler_get(c *gin.Context) {
	token := c.DefaultQuery("token", "invalid")
	if token == "invalid" {
		c.Redirect(http.StatusTemporaryRedirect, "/list")
		return
	} else {
		if val, ok := ServersMap[token]; ok {
			c.HTML(http.StatusOK, "index.html", gin.H{"token": token, "ip": val.Get("Ip"), "port": val.Get("Port")})
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/list")
			return
		}
	}
}

func LoginHandler_get(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginHandler_post(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./ACCOUNT.sqlite")
	checkErr(err)
	var user User
	err = c.ShouldBind(&user)
	checkErr(err)
	var passwordInDb string
	rows, err := db.Query(fmt.Sprintf("SELECT PASSWORD FROM USER WHERE NAME = %q", user.Name))
	for rows.Next() {
		err = rows.Scan(&passwordInDb)
		checkErr(err)
		break
	}
	rows.Close()
	db.Close()
	if user.Password != passwordInDb {
		//log.Println("login failed")
		c.String(http.StatusOK, "failed")
	} else {
		//c.SetCookie(CookieName, CookieValue, 0, "/", "", false, true)
		session := sessions.Default(c)
		session.Set("login", "true")
		session.Save()
		c.String(http.StatusOK, "ok")
	}
}

func ProcessHandler_get(c *gin.Context) {
	token := c.DefaultQuery("token", "invalid")
	if token == "invalid" {
		c.Redirect(http.StatusTemporaryRedirect, "/list")
		return
	} else {
		if val, ok := ServersMap[token]; ok {
			c.HTML(http.StatusOK, "processes.html", gin.H{"token": token, "ip": val.Get("Ip"), "port": val.Get("Port")})
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/list")
			return
		}
	}
}

func PathHandler_get(c *gin.Context) {
	token := c.DefaultQuery("token", "invalid")
	var val gjson.Result
	if token == "invalid" {
		c.Redirect(http.StatusTemporaryRedirect, "/list")
		return
	} else {
		if v, ok := ServersMap[token]; ok {
			val = v
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/list")
			return
		}
	}

	path := c.DefaultQuery("path", "/")
	resp, err := http.Get("http://" + val.Get("Ip").Str + ":" + val.Get("Port").Str + "/path?token=" + token + "&path=" + path)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/list")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/list")
		return
	}
	var pathinfo PathInfo
	json.Unmarshal(body, &pathinfo)
	c.HTML(http.StatusOK, "path.html", gin.H{"token": token, "ip": val.Get("Ip"), "port": val.Get("Port"), "header": gjson.Get(string(body), "path").Str, "writable": gjson.Get(string(body), "writable").Bool(), "dirs": pathinfo.Dirs, "files": pathinfo.Files})
}

func AboutHandler_get(c *gin.Context) {
	token := c.DefaultQuery("token", "invalid")
	if token == "invalid" {
		c.Redirect(http.StatusTemporaryRedirect, "/list")
		return
	} else {
		if val, ok := ServersMap[token]; ok {
			c.HTML(http.StatusOK, "about.html", gin.H{"token": token, "ip": val.Get("Ip"), "port": val.Get("Port")})
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/list")
			return
		}
	}
}

func AuthorsHandler_get(c *gin.Context) {
	token := c.DefaultQuery("token", "invalid")
	if token == "invalid" {
		c.Redirect(http.StatusTemporaryRedirect, "/list")
		return
	} else {
		if val, ok := ServersMap[token]; ok {
			c.HTML(http.StatusOK, "authors.html", gin.H{"token": token, "ip": val.Get("Ip"), "port": val.Get("Port")})
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/list")
			return
		}
	}
}

func ListServerHandler_post(c *gin.Context) {
	c.JSON(http.StatusOK, ServersString)
}

func AddServerHandler_post(c *gin.Context) {
	token := c.PostForm("token")
	server := c.PostForm("server")
	result := gjson.Get(server, "Servers")
	ServersMap[token] = result
	j := make(map[string][]Server)
	j["Servers"] = make([]Server, 0)
	for _, v := range ServersMap {
		var test2 Server
		test2.Ip = v.Get("Ip").Str
		test2.Token = v.Get("Token").Str
		test2.Port = v.Get("Port").Str
		test2.Username = v.Get("Username").Str
		test2.Password = v.Get("Password").Str
		j["Servers"] = append(j["Servers"], test2)
	}
	data, _ := json.Marshal(j)
	ioutil.WriteFile("servers.json", data, 0644)
	ServersString = string(data)
	c.String(http.StatusOK, "ok")
}

func DeleteServerHandler_post(c *gin.Context) {
	id := c.PostForm("token")
	delete(ServersMap, id)
	j := make(map[string][]Server)
	j["Servers"] = make([]Server, 0)
	for _, v := range ServersMap {
		var test2 Server
		test2.Ip = v.Get("Ip").Str
		test2.Token = v.Get("Token").Str
		test2.Port = v.Get("Port").Str
		test2.Username = v.Get("Username").Str
		test2.Password = v.Get("Password").Str
		j["Servers"] = append(j["Servers"], test2)
	}
	data, _ := json.Marshal(j)
	ioutil.WriteFile("servers.json", data, 0644)
	ServersString = string(data)
	c.String(http.StatusOK, "ok")
}

func ListHandler_get(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{})
}

func ChangeTokenHandler_post(c *gin.Context) {
	pre := c.PostForm("pre")
	now := c.PostForm("now")
	if pre != "" && now != "" && pre != now {
		ServersString = strings.Replace(ServersString, pre, now, 1)

		result := gjson.Get(ServersString, "Servers")
		ServersMap = make(map[string]gjson.Result)
		for _, val := range result.Array() {
			ServersMap[val.Get("Token").Str] = val
		}

		log.Println(ServersMap)

		ioutil.WriteFile("servers.json", []byte(ServersString), 0644)
		c.String(http.StatusOK, "ok")
		return
	}

	c.String(http.StatusOK, "no change")
}

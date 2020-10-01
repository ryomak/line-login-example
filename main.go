package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"math/rand"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

const (
	redirectURL = "http://localhost:8080/redirect"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Random":      xid.New().String(),
			"ChannelID":   "1655032271",
			"RedirectURL": redirectURL,
		})
	})

	r.GET("/redirect", func(c *gin.Context) {
		// access token
		code := c.Query("code")
		values := url.Values{}
		values.Set("grant_type", "authorization_code")
		values.Add("code", code)
		values.Add("redirect_uri", redirectURL)
		values.Add("client_id", "1655032271")
		values.Add("client_secret", "")
		req, _ := http.NewRequest("POST", "https://api.line.me/oauth2/v2.1/token", strings.NewReader(values.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			c.HTML(http.StatusOK, "redirect.tmpl", gin.H{
				"Error": err.Error(),
				"Res":   "",
			})
			return
		}
		defer resp.Body.Close()
		var params struct {
			AccessToken string `json:"access_token"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&params); err != nil {
			c.HTML(http.StatusOK, "redirect.tmpl", gin.H{
				"Error": err.Error(),
				"Res":   "",
			})
			return
		}
		fmt.Println(params.AccessToken)
		// user 情報
		req, _ = http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", params.AccessToken))
		resp, err = client.Do(req)
		if err != nil {
			c.HTML(http.StatusOK, "redirect.tmpl", gin.H{
				"Error": err.Error(),
				"Res":   "",
			})
			return
		}
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)

		c.HTML(http.StatusOK, "redirect.tmpl", gin.H{
			"Error": "",
			"Res":   string(b),
		})
	})
	r.Run(":8080")
}

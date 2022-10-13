package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"ng-ups/config"
)

// 获取get请求数据
func Get(url string) (string, error) {
	if config.Config.GetBool("debug") {
		fmt.Println(url)
	}
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", nil
	}
	body, err := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func Filter(content string, filter string) string {
	strList := strings.Split(content, "\n")
	for _, s := range strList {
		if strings.Contains(s, filter) {
			return s
		}
	}

	return ""
}

func GetP(query map[string]string, filter string) string {
	url := config.Config.GetString("url")
	zone := config.Config.GetString("zone")
	var queryString []string
	query["upstream"] = zone
	for key, value := range query {
		queryString = append(queryString, key+"="+value)
	}
	fullUrl := url + "?" + strings.Join(queryString, "&")

	content, _ := Get(fullUrl)
	if content != "" && filter != "" {
		content = Filter(content, filter)
	}

	return content
}

func GetServer(server string) string {
	return GetP(map[string]string{"verbose": ""}, server)
}

func Debug(content string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), content)
}
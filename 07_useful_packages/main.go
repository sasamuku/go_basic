package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"time"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
}

func main() {
	// fmt

	// Printの違い
	// fmt.Print() が基本
	// {Prefix}Print{Suffix}
	// Prefix: F
	//

	// time
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339)) // RFC3339: PostgreSQL で標準的に使われる時間表記

	// regex (正規表現)
	// 正規表現のチェックを単発で行いたいとき
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	match2, _ := regexp.MatchString("a([a-z]+)e", "appl0e")
	fmt.Println(match, match2)

	// 正規表現のチェックを複数回行いたいとき
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fs2 := r2.FindString("/delete/test")
	fmt.Println(fs, "...", fs2)

	// 正規表現のチェックを複数回行いたいとき、かつ、部分文字列を取り出したいとき
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/local")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	// fss = r2.FindStringSubmatch("/delete/test")
	// fmt.Println(fss, fss[0], fss[1], fss[2]) // 値を取得できないのでエラーとなる

	// sort
	i := []int{2, 5, 1, 3, 0}
	s := []string{"Mike", "Alice", "Bob", "Alan"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Bob", 21},
		{"Judy", 19},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	// sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)

	// iota
	// context
	// goroutineのタイムアウト設定ができる

	// ioutil
	// ファイルの読み書き

	// http
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	base, _ := url.Parse("http://example.com")
	ref, _ := url.Parse("/test")
	endpoint := base.ResolveReference(ref).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("header", "value")

	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// json.Unmarshal
	j := []byte(`{"name":"Mike", "age":20, "nicknames":["a","b","c"]}`)
	var p2 Person
	if err2 := json.Unmarshal(j, &p2); err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(p2.Name, p2.Age, p2.Nicknames)

	v, _ := json.Marshal(p2)
	fmt.Println(string(v))

	// hmac
	// ハッシュ生成してAPI認証
	test := []byte("Hello")
	fmt.Println(test)
}

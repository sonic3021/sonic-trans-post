package imp

import (
	"net/http"
	"fmt"
	"net/url"
)

func hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"hello\n")
}

func headers(w http.ResponseWriter, req *http.Request){
	for name, headers:=range req.Header{
		for _, h := range headers{
			fmt.Fprintf(w, "%v, %v\n", name, h)
		}
	}
}

func transPost(w http.ResponseWriter, req *http.Request){
	var purl="aa"
	var msg="bb"

	postHttp(purl ,msg )

}


func postHttp(purl string, msg string ){
	// postBody, _ := json.Marshal(map[string]string{
	// 	"name":  "Toby",
	// 	"email": "Toby@example.com",
	// })
	data := url.Values{
		"to":       {},
		"msg":      {},
	}
	resp, err := http.PostForm("", data)

	if err != nil {
		fmt.Println(err)
	}
	//var res map[string]interface{}

	//json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(resp.Body)

	//fmt.Println(res["form"])
}

func ServeHttp(){
	fmt.Println("start http server...")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/trans-post", transPost)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println(err)
	}
}

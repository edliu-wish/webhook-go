package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	max_message_length = 5
)

var list *List

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Webhook Demo</h1>")
	fmt.Fprintf(w, "<h2>server is ready to receive message</h2>")
	fmt.Fprintf(w, "<p>you can post message to this url <a href=\"http://localhost:8080/alert\">http://localhost:8080/alert</a> with post json body, and you will see the messages here</p>")
	fmt.Fprintf(w, "<p>you can also get the latest message from this url: <a href=\"http://localhost:8080/latest\">http://localhost:8080/latest</a></p>")
	fmt.Fprintf(w, "<h2>Received Messages:</h2>")
	for i := uint(0); i < (*list).size; i++ {
		node := (*list).Get(i)
		fmt.Fprintf(w, "<b>content: </b><code>"+(*node).message+"</code> <b>, received time: </b><code>"+((*node).timestamp).String()+"</code><br><br/>")
	}

}

func latestHandle(w http.ResponseWriter, r *http.Request) {
	if (*list).size == 0 {
		fmt.Fprintf(w, "no latest message")
		return
	}
	node := (*list).tail
	fmt.Fprintf(w, "<b>content: </b><code>"+(*node).message+"</code> <b>, received time: </b><code>"+((*node).timestamp).String()+"</code><br><br/>")
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	node := new(WebhookMessage)
	node.message = string(body)
	node.timestamp = time.Now()
	(*list).Append(node)
	if (*list).size > max_message_length {
		(*list).RemoveHead()
	}
	fmt.Fprintf(w, "ok")
}

func main() {
	list = new(List)
	http.HandleFunc("/", listHandler) //	设置访问路由
	http.HandleFunc("/alert", pushHandler)
	http.HandleFunc("/latest", latestHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

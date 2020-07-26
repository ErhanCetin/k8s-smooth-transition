package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-stomp/stomp"
	"gopkg.in/mgo.v2"
	"log"
)

var serverAddr = flag.String("server", "localhost:61613", "STOMP server endpoint")

// TODO :
// TODO change news-docker que name also in Pythong script
var queueName = flag.String("queue", "/queue/news-queue", "Destination queue")
var stop = make(chan bool)

// Reference https://godoc.org/github.com/barthr/newsapi
type Source struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type NewsMessage struct {
	Status       string `json:"status,omitempty"`
	TotalResults int32  `json:"totalResults,omitempty"`
	Articles     []struct {
		Source      Source `json:"source,omitempty"`
		Author      string `json:"author,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Url         string `json:"url,omitempty"`
		UrlToImage  string `json:"urlToImage,omitempty"`
		PublishedAt string `json:"publishedAt,omitempty`
		Content     string `json:"content,omitempty"`
	} `json:"articles,omitempty"`
}

func main() {
	startConsumer()
}

func startConsumer() {
	subscribed := make(chan bool)

	go consumerNews(subscribed)

	// wait until we know the receiver has subscribed

	<-stop
}

func consumerNews(subscribed chan bool) {
	defer func() {
		stop <- false
	}()

	conn, err := stomp.Dial("tcp", *serverAddr)

	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}

	sub, err := conn.Subscribe(*queueName, stomp.AckAuto)
	if err != nil {
		println("cannot subscribe to", *queueName, err.Error())
		return
	}
	//close(subscribed)

	msg := <-sub.C
	actualText := string(msg.Body)
	println("Actual:", actualText)
	println("The message is received.")

	// REFERENCE : https://itnext.io/welcome-to-just-enough-go-7dbef7e30188
	var dat NewsMessage
	if err := json.Unmarshal(msg.Body, &dat); err != nil {
		panic(err)
	}

	/* second way to decode news message
	   jsonDataReader := strings.NewReader(actualText)
	   decoder := json.NewDecoder(jsonDataReader)

	   if err := decoder.Decode(&dat); err != nil {
	      panic(err)
	   }*/
	fmt.Println(dat)
	startConsumer()


}
func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("news-db").C("article")

	// for details of index : https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
	index := mgo.Index{
		Key:        []string{"source"},
		Unique:     false,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func addNews(s *mgo.Session, class NewsMessage) {
	session := s.Copy()
	defer session.Close()

	/*var book Book
	  decoder := json.NewDecoder(r.Body)
	  err := decoder.Decode(&book)
	  if err != nil {
	     ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
	     return
	  }*/

	c := session.DB("news-db").C("article")

	for i, s := range class.Articles {
		fmt.Println(i, s)
		if err := c.Insert(s); err != nil {
			if mgo.IsDup(err) {
				log.Println("News with this source already exists: ", err)
				return
			}
			log.Println("Failed insert News: ", err)
		}
	}
	return
}

func newsBySource(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()
	c := session.DB("news-db").C("article")

	/* // get by soure .. but It does not work .
	   sourcenews := Source{
	         Id:   "",
	         Name: "Lifehacker.com"}



	      var news NewsMessage
	      err := c.Find(bson.M{"source": sourcenews}).One(&news)
	      if err != nil {
	         log.Println("Failed find news: ", err)
	         return
	      }*/

	// get all news
	var newsList []NewsMessage
	err := c.Find(nil).All(&newsList)
	if err != nil {
		log.Println("Failed find news: ", err)
		return
	} else {
		fmt.Println("Results All: ", newsList)
	}


}

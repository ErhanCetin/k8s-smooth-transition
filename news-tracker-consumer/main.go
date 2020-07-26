package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-stomp/stomp"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

var stop = make(chan bool)

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

	var activemqHost = "localhost"
	var activemqPort = "61613"

	if hostName := os.Getenv("ACTIVEMQ-HOST"); hostName != "" {
		activemqHost = hostName
	}
	if hostPort := os.Getenv("ACTIVEMQ-PORT"); hostPort != "" {
		activemqPort = hostPort
	}
	var serverAddr = flag.String("server", activemqHost+":"+activemqPort, "STOMP server endpoint")

	var activemqQueueName = "news-queue"
	if queueName := os.Getenv("ACTIVEMQ-QUEUE-NAME"); queueName != "" {
		activemqQueueName = queueName
	}

	var queueNameFlag = flag.String("queue", "/queue/"+activemqQueueName, "Destination queue")
	subscribed := make(chan bool)

	go consumerNews(subscribed, serverAddr, queueNameFlag)

	// wait until we know the receiver has subscribed
	<-subscribed
	<-stop
}

func consumerNews(subscribed chan bool, serverAddr *string, queueName *string) {
	defer func() {
		stop <- true
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
	close(subscribed)
	msg := <-sub.C
	processRecievedMessage(msg)
}

func processRecievedMessage(msg *stomp.Message) {

	actualText := string(msg.Body)
	println("Actual:", actualText)
	println("The message is received.")

	// REFERENCE : https://itnext.io/welcome-to-just-enough-go-7dbef7e30188
	var dat NewsMessage
	if err := json.Unmarshal(msg.Body, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	var mongoDbHost = "localhost"

	if host := os.Getenv("MONGODB-HOST"); host != "" {
		mongoDbHost = host
	}
	session, err := mgo.Dial(mongoDbHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var mongoDb = "news-db"
	var mongoCollection = "article"

	if dbName := os.Getenv("MONGODB-NAME"); dbName != "" {
		mongoDb = dbName
	}

	if collection := os.Getenv("MONGODB-COLLECTION"); collection != "" {
		mongoCollection = collection
	}

	createIndex(session, mongoDb, mongoCollection)
	//newsBySource(session)
	addNews(session, dat, mongoDb, mongoCollection)
}

func createIndex(s *mgo.Session, db string, collection string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(db).C(collection)

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

func addNews(s *mgo.Session, class NewsMessage, db string, collection string) {
	session := s.Copy()
	defer session.Close()

	//TODO parameterize mongo db and collection
	c := session.DB(db).C(collection)

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

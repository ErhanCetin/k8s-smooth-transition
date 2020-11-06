package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"strconv"

	stomper "github.com/russmack/stompingophers"
)

var (
	printer   chan stomper.ServerFrame
	client    stomper.Client
	queueIP   = "localhost"
	queuePort = 61613
	queueName = "/queue/news-queue"
)

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
/*
  reference : https://github.com/russmack/stompingophers/blob/master/stompingophers.go
 */
func main() {

	if hostName := os.Getenv("ACTIVEMQ-HOST"); hostName != "" {
		queueIP = hostName
	}
	if hostPort := os.Getenv("ACTIVEMQ-PORT"); hostPort != "" {
		queuePort, _ = strconv.Atoi(hostPort)
	}

	if queueNameEnv := os.Getenv("ACTIVEMQ-QUEUE-NAME"); queueNameEnv != "" {
		queueName = queueNameEnv
	}

	printer = make(chan stomper.ServerFrame)

	go func() {
		for {
			msg := <-printer
			_ = msg
		}
	}()
	c := connect()
	sub := subscribe(c)
	consumer(c, sub)
}

func connect() *stomper.Client {

	var err error
	conn, err := stomper.NewConnection(queueIP, queuePort)
	if err != nil {
		log.Fatal(err)
	}

	options := stomper.Options{
		HeartBeat: &stomper.HeartBeat{
			SendInterval: 400000,
			RecvTimeout:  400000,
		},
	}

	client, resp, err := stomper.Connect(conn, &options)
	if err != nil {
		log.Fatal("failed connecting: " + err.Error())
	}

	f, err := stomper.ParseResponse(resp)
	if err != nil {
		log.Fatal("failed parsing connect response:", err)
	}
	fmt.Printf("Conneced: %+v\n", f)

	return &client
}

func subscribe(client *stomper.Client) stomper.Subscription {
	fmt.Println("Subscribing to queue...\n")

	sub, resp, err := client.Subscribe(queueName, "mysubrcpt", stomper.AckModeClient)
	if err != nil {
		log.Fatal("failed sending: " + err.Error())
	}

	msg, err := stomper.ParseResponse(resp)
	if err != nil {
		fmt.Println("failed parsing subscribe response:", err)
	}
	fmt.Printf("Parsed subscribe response:\n%s\n", string(msg.Headers["receipt-id"]))

	return sub
}

func consumer(client *stomper.Client, sub stomper.Subscription) {
	defer client.Disconnect()

	recvChan, errChan := client.Receive()

	for {
		select {
		case err := <-errChan:
			log.Printf("This is unfortunate, but the show must go on.  (err: %s)", err)
		case s := <-recvChan:
			f, err := stomper.ParseResponse(s)
			if err != nil {
				fmt.Println("response parse err:", err)
				continue
			}
			processRecievedMessage(f.Body)

			// TODO chnage =! dont forget
			if sub.AckMode == stomper.AckModeAuto {
				msgAckID := string(f.Headers["ack"])
				err = client.Ack(msgAckID, "", "")
				if err != nil {
					fmt.Println("failed sending ack:", err)
					continue
				}
			}
			printer <- f
		}
	}

}

func processRecievedMessage(queueMessageBody []byte) {

	actualText := string(queueMessageBody)
	println("Actual:", actualText)
	println("The message is received.")

	// REFERENCE : https://itnext.io/welcome-to-just-enough-go-7dbef7e30188
	var dat NewsMessage
	err := json.Unmarshal(bytes.Trim(queueMessageBody, "\x00"), &dat)
	if err == nil {

		var mongoDbHost = "127.0.0.1:27017"
		var mongoDb = "news-db"
		var mongoCollection = "article"

		if host := os.Getenv("MONGODB-HOST"); host != "" {
			mongoDbHost = host
		}
		if dbName := os.Getenv("MONGODB-NAME"); dbName != "" {
			mongoDb = dbName
		}
		if collection := os.Getenv("MONGODB-COLLECTION"); collection != "" {
			mongoCollection = collection
		}

		session, err := mgo.Dial(mongoDbHost)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)

		createIndex(session, mongoDb, mongoCollection)
		addNews(session, dat, mongoDb, mongoCollection)
		fmt.Println(dat)
	} else {
		fmt.Println(err)
	}

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

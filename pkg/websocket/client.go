package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "webapi.wiegand.com.cn:61081", "http service address")

// {
//        "SN": 223289803,
//        "1\u53f7\u95e8": "\u5f00",
//        "2\u53f7\u95e8": "\u5f00",
//        "\u5237\u65b0\u65f6\u95f4": "2022-06-20 18:50:53"
//    }
//       "SN": 223289803,
//       "1号门": "开",
//      "2号门": "开",
//        "刷新时间": "2022-06-20 18:50:53"

type ASCII1 struct {
	SN        int    `json:"SN"`
	OneDoor   string `json:"1号门"`
	TwoDoor   string `json:"2号门"`
	FlushTime string `json:"刷新时间"`
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var ret []ASCII1
			err = json.Unmarshal(message, &ret)
			log.Printf("%v", ret)
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		// case t := <-ticker.C:
		// 	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

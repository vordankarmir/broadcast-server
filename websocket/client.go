package websocket

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"

	socketio "github.com/maldikhan/go.socket.io/socket.io/v5/client"
)

func ConnectClient() *socketio.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := socketio.NewClient(
		socketio.WithRawURL("http://127.0.0.1:8000/"),
	)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	client.On("message", func(data []byte) {
		fmt.Printf("Received message: %s\n", string(data))
	})

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}

	if err := client.Emit("message", "I just connected"); err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	<-ctx.Done()

	return client
}

func EmitEvent(client *socketio.Client, message string) {
	err := client.Emit("message", message)

	if err != nil {
		log.Fatalf("Error sending message")
	}
}

func ConnectWebsocketClient() {

	var addr = flag.String("addr", "http://127.0.0.1:8000", "http service address")

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
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
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
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

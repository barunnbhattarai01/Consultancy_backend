package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 60 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan Event
	chatroom   string
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		//egress is used to avoid concurrent writes on the websocket connrction
		//it needed because if two goroutine try to write message at same time ,it can panic and corrupt the mesage
		egress: make(chan Event),
	}
}

// read messages
func (c *Client) readMessages() {
	defer func() {
		//cleanup connection
		c.manager.removeClient(c)
	}()

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
	c.connection.SetReadLimit(512)
	c.connection.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message:%v", err)
			}
			break
		}
		//skip the invalid json
		if !json.Valid(payload) {
			log.Println("Invalid JSON received, skipping:", string(payload))
			continue
		}

		var request Event

		if err := json.Unmarshal(payload, &request); err != nil {
			log.Print("error marshelling event", err)
			break
		}

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println("error handling message:", err)
		}
	}
}

//func to write message

func (c *Client) writeMessage() {
	defer func() {
		c.manager.removeClient(c)
	}()

	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("conection closed:", err)
				}
				return
			}
			data, err := json.Marshal(message)
			if err != nil {
				log.Print(err)
				return
			}
			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("failed to send message:", err)

			}
			log.Println("message sent")

		case <-ticker.C:
			log.Println("ping")

			//send ping to client
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Println("write message error:", err)
				return
			}
		}
	}

}

func (c *Client) pongHandler(pongMsg string) error {
	log.Print("pong")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

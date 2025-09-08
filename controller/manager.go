package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	Clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		Clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandler()
	return m
}

func (m *Manager) setupEventHandler() {
	m.handlers[EventSendMessage] = SendMessage
	m.handlers[EventChatRoom] = ChatRoomhandler
}

func ChatRoomhandler(event Event, c *Client) error {
	var ChangeroomEvent ChangeroomEvent
	if err := json.Unmarshal(event.Payload, &ChangeroomEvent); err != nil {
		return fmt.Errorf("bad payload is request :%v", err)
	}
	c.chatroom = ChangeroomEvent.Name
	return nil
}

func SendMessage(event Event, c *Client) error {
	var chatevent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request€€:%v", err)
	}

	var broadMessage NewMessage

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From

	data, err := json.Marshal(broadMessage)

	if err != nil {
		return fmt.Errorf("failed to marshell:%v ", err)
	}

	outgoingEvent := Event{
		Payload: data,
		Type:    EventNewMessage,
	}

	for client := range c.manager.Clients {
		if client.chatroom == c.chatroom {
			client.egress <- outgoingEvent

		}
	}

	return nil
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	//check the event type
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Printf("new Connectionn")
	websocketUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	//upgarde regular http connection into websocket
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("error in connection of websocket")
		return
	}
	client := NewClient(conn, m)
	m.addClient(client)

	//start clinet process using goroutine
	go client.readMessages()
	go client.writeMessage()
}

// addign helping func fro adding  clinet
func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.Clients[client] = true

}

// func to removing client
func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Clients[client]; ok {
		err := client.connection.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("error sending close:", err)
		}
		client.connection.Close()
		delete(m.Clients, client)
	}

}

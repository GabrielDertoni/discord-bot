package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// GatewayManager - Gateway manager
type GatewayManager struct {
	Connection        *websocket.Conn
	HeartbeatInterval int
	Events            chan Payload
	Done              chan struct{}
	heartbeatACK      chan struct{}
}

// NewGatewayManager - Creates a gateway manager
func NewGatewayManager() *GatewayManager {
	manager := GatewayManager{}
	return &manager
}

// OpenConnection - Connects the manager to the gateway with specific intents.
func (manager *GatewayManager) OpenConnection(token string, intents int) (err error) {
	connection, _, err := websocket.DefaultDialer.Dial(GatewayURL, nil)
	if err != nil {
		return
	}
	manager.Connection = connection
	manager.Events = make(chan Payload, 100)
	manager.Done = make(chan struct{})
	manager.heartbeatACK = make(chan struct{}, 1)

	go manager.eventLoop()

	heartbeatInterval, err := waitForHello(manager.Events)
	if err != nil {
		return
	}
	manager.HeartbeatInterval = heartbeatInterval
	go manager.heartbeat()

	err = manager.SendIdentify(token, intents)
	if err != nil {
		return
	}
	return nil
}

// SendIdentify - Sends identify message through manager and sets presence to online
func (manager *GatewayManager) SendIdentify(token string, intents int) (err error) {
	identifyData := Identify{
		Token:   token,
		Intents: intents,
		Properties: IdentifyProperties{
			OS:      "linux",
			Browser: "go-discord",
			Device:  "go-discord",
		},
		Presence: PresenceUpdate{
			Status: StatusOnline,
			AFK:    false,
		},
	}
	jsonIdentify, err := json.Marshal(identifyData)
	if err != nil {
		return
	}
	err = manager.Connection.WriteJSON(Payload{
		Op:   OpCodeIdentify,
		Data: jsonIdentify,
	})
	return
}

// ChangeToOfflineStatus - Makes bot offline
func (manager *GatewayManager) ChangeToOfflineStatus() (err error) {
	err = manager.UpdatePresence(PresenceUpdate{
		Status: StatusOffline,
		AFK:    true,
	})
	return
}

// ChangeToOnlineStatus - Makes bot online
func (manager *GatewayManager) ChangeToOnlineStatus() {
	manager.UpdatePresence(PresenceUpdate{
		Status: StatusOnline,
		AFK:    false,
	})
	fmt.Println("Status is offline.")
}

func (manager *GatewayManager) UpdatePresence(data PresenceUpdate) (err error) {
	jsonStatus, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = manager.Connection.WriteJSON(Payload{
		Op:   OpCodePresenceUpdate,
		Data: jsonStatus,
	})
	return
}

// CloseConnection - Closes the websocket connection
func (manager *GatewayManager) CloseConnection() {
	fmt.Println("Closing connection...")
	manager.ChangeToOfflineStatus()
	close(manager.Done)
	openEvents, openHeartbeat := true, true
	// Wait for manager.Events and manager.heartbeatACK to close
	for openEvents || openHeartbeat {
		select {
		case _, openEvents = <-manager.Events:
		case _, openHeartbeat = <-manager.heartbeatACK:
		}
	}
	manager.Connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	manager.Connection.Close()
}

// Heartbeat - Routine that sends heartbeats every manager.HeartbeatInterval
func (manager *GatewayManager) heartbeat() {
	fmt.Printf("Sending heartbeat every %d milliseconds\n", manager.HeartbeatInterval)
	ticker := time.NewTicker(time.Duration(manager.HeartbeatInterval) * time.Millisecond)
	ACKTimer := time.NewTimer(0)
	ACKTimer.Stop()
	for {
		select {
		case <-ticker.C:
			heartbeat := Payload{Op: OpCodeHeartbeat}
			err := manager.Connection.WriteJSON(heartbeat)
			if err != nil {
				panic(err)
			}
			ACKTimer.Reset(1000 * time.Millisecond)

		case <-manager.heartbeatACK:
			ACKTimer.Stop()

		case <-ACKTimer.C:
			log.Println("Heartbeat ACK timeout.")
			// panic(errors.New("Heartbeat ACK timeout."))

		case <-manager.Done:
			ticker.Stop()
			ACKTimer.Stop()
			close(manager.heartbeatACK)
			return
		}
	}
}

func (manager *GatewayManager) eventLoop() {
	messageChan := make(chan []byte, 1)
	go websocketMessageReader(manager.Connection, messageChan, manager.Done)
	for {
		select {
		case <-manager.Done:
			fmt.Println("Exiting event loop")
			close(manager.Events)
			return
		case message := <-messageChan:
			var payload Payload
			err := json.Unmarshal(message, &payload)
			if err != nil {
				log.Println("Event loop error:", err)
				log.Println("Tried to unmarshal:", string(message))
				return
			}
			if payload.Op == OpCodeHeartbeatACK {
				manager.heartbeatACK <- struct{}{}
			} else {
				manager.Events <- payload
			}
		}
	}
}

// WaitForHello - Waits for a hello message (op code 10) from websocket
func waitForHello(eventChan chan Payload) (interval int, err error) {
	payload := <-eventChan
	if payload.Op != 10 {
		err = errors.New("Expected op code 10")
		return
	}
	data := Hello{}
	err = json.Unmarshal(payload.Data, &data)
	if err != nil {
		return
	}
	interval = data.HeartbeatInterval
	return
}

func websocketMessageReader(connection *websocket.Conn, channel chan<- []byte, done chan<- struct{}) {
	for {
		messageType, message, err := connection.ReadMessage()
		if messageType == -1 {
			close(channel)
			return
		}
		if err != nil {
			log.Println("websocketMessageReader error:", err)
		}
		channel <- message
	}
}

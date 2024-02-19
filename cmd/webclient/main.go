package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

func main() {
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt)

	u := "ws://localhost:8080/ws"
	log.Println("Connecting to", u)

	// Start the web client
	c, _, err := websocket.DefaultDialer.Dial(u, nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	fmt.Println("Starting up")
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = pc.Close(); err != nil {
			fmt.Printf("cannot close peerConnection: %v\n", err)
		}
	}()

	// Create a datachannel with label 'data'
	dataChannel, err := pc.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}

	// Register text message handling
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Message from DataChannel '%s': '%s'\n", dataChannel.Label(), string(msg.Data))
	})

	offer, err := pc.CreateOffer(nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := pc.SetLocalDescription(offer); err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	// read messages from the server
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

	// write messages to the server

	message, err := json.Marshal(offer)
	if err != nil {
		log.Println("json:", err)
		return
	}

	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case <-interupt:
			log.Println("Stop web client")
			// Cleanly close the connection by sending a close message and then
			if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
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

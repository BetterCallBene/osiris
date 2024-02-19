package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

var upgrader = websocket.Upgrader{} // use default options

func main() {

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
			log.Printf("cannot close peerConnection: %v\n", err)
		}
	}()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		defer conn.Close()

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		err,

		// offer := webrtc.SessionDescription{}

		// // unmarshal the offer
		// err = json.Unmarshal(message, &offer)
		// if err != nil {
		// 	log.Println("unmarshal:", err)
		// 	return
		// }

		// if err := pc.SetRemoteDescription(offer); err != nil {
		// 	log.Printf("cannot set remote description: %v\n", err)
		// 	return
		// }

		// answer, err := pc.CreateAnswer(nil)

		// if err != nil {
		// 	log.Println("create answer:", err)
		// 	return
		// }

		// payload, err := json.Marshal(answer)
		// if err != nil {
		// 	log.Println("json:", err)
		// 	return
		// }

		// err = conn.WriteMessage(websocket.TextMessage, payload)
		// if err != nil {
		// 	log.Println("write:", err)
		// 	return
		// }

		// if err := pc.SetLocalDescription(answer); err != nil {
		// 	log.Println("set local description:", err)
		// 	return
		// }

	})

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	pc.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			fmt.Println("Peer Connection has gone to failed exiting")
			//os.Exit(0)
		}

		if s == webrtc.PeerConnectionStateClosed {
			// PeerConnection was explicitly closed. This usually happens from a DTLS CloseNotify
			fmt.Println("Peer Connection has gone to closed exiting")
			//os.Exit(0)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

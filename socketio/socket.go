package socketio

import (
    "log"
    socketio "github.com/googollee/go-socket.io"
)


func SocketIO() (*socketio.Server) {
	server := socketio.NewServer(nil)
	if server == nil {
		log.Fatal("Failed to create socket.io server")
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason , s.ID())
	})
	return server
}

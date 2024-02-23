package utils

import "github.com/gorilla/websocket"

func FindConnIndex(slice []*websocket.Conn, conn *websocket.Conn) int {
	for index, connection := range slice {
		if connection == conn {
			return index
		}
	}
	return -1
}

func RemoveConn(slice []*websocket.Conn, index int) []*websocket.Conn {
	return append(slice[:index], slice[index+1:]...)
}

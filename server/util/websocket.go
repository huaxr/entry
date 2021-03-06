package util

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mijia/sweb/log"

	"github.com/laincloud/entry/server/message"
)

func SendCloseMessage(conn *websocket.Conn, content []byte, msgMarshaller Marshaler, writeLock *sync.Mutex) {
	closeMsg := &message.ResponseMessage{
		MsgType: message.ResponseMessage_CLOSE,
		Content: content,
	}
	if closeData, err := msgMarshaller(closeMsg); err != nil {
		log.Errorf("Marshal close message failed: %s", err.Error())
	} else {
		writeLock.Lock()
		conn.WriteMessage(websocket.BinaryMessage, closeData)
		writeLock.Unlock()
	}
}

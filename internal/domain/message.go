package domain

import "time"

// Message — доменное сообщение (до шифрования)
type Message struct {
	ID        string
	From      PeerID
	To        PeerID
	Text      string
	Timestamp time.Time
}

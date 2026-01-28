package domain

import "crypto/sha256"

// PeerID — короткий идентификатор узла
type PeerID string

// Peer — участник сети
type Peer struct {
	ID        PeerID
	PublicKey []byte
	Address   string
}

// NewPeerID генерирует PeerID из публичного ключа
func NewPeerID(publicKey []byte) PeerID {
	hash := sha256.Sum256(publicKey)
	return PeerID(hash[:16])
}

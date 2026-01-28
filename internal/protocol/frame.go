package protocol

import "github.com/Asxatillo/meshtalk/internal/domain"

// FrameType — тип кадра протокола
type FrameType uint8

const (
	FrameDiscovery FrameType = 1
	FrameHandshake FrameType = 2
	FrameMessage   FrameType = 3
	FramePing      FrameType = 4
	FrameError     FrameType = 255
)

// Frame — базовая единица протокола
type Frame struct {
	Version uint8         `json:"version"`
	Type    FrameType     `json:"type"`
	From    domain.PeerID `json:"from"`
	Payload []byte        `json:"payload"`
}

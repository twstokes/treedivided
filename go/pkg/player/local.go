package player

// LocalPlayer conforms to the Player interface
type LocalPlayer struct{}

// NewLocalPlayer creates a new player that runs locally on the machine
func NewLocalPlayer() *LocalPlayer {
	return &LocalPlayer{}
}

package tree

// TeamID is a unique identifier for each team
type TeamID uint8

const (
	// TeamA is the first team
	TeamA TeamID = 1
	// TeamB is the second team
	TeamB TeamID = 2
)

// Team is a sports team
type Team struct {
	ID       TeamID
	Name     string
	Colors   TeamColors
	SongPath string
}

// TeamColors are the primary and secondary colors for a team
type TeamColors struct {
	Primary   Color
	Secondary Color
}

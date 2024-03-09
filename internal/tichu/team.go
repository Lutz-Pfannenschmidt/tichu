package tichu

type Team struct {
	Players [2]*Player
	Game    *Game
}

func NewTeam(game *Game) *Team {
	return &Team{Players: [2]*Player{}, Game: game}
}

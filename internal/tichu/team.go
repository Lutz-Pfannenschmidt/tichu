package tichu

type Team struct {
	Players [2]*Player
	Game    *Game
	ID      int
}

func NewTeam(game *Game, id int, playerName1, playerName2 string) *Team {
	res := &Team{Players: [2]*Player{}, Game: game, ID: id}
	res.Players[0] = NewPlayer(playerName1, res, id*2)
	res.Players[1] = NewPlayer(playerName2, res, id*2+1)
	return res
}

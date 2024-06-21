package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

type Die struct {
	Color int
	Lose  int
	Win   int
}

const (
	Brain     = "brain🧠"
	Shotgun   = "shotgun⌖"
	Footprint = "footprint👣"
)

var (
	greenDie  = Die{Color: 32, Lose: 1, Win: 3}
	yellowDie = Die{Color: 33, Lose: 2, Win: 4}
	redDie    = Die{Color: 31, Lose: 3, Win: 5}
	emptyDie  = Die{}
)

type Brains struct {
	Points int
	Dice   []Die
}

type Player struct {
	ID    int
	Score int
}

func (die Die) GetColoredDie(roll int) string {
	const colorStr = "\x1b[%dm%s\x1b[0m" //ANSI color codes for the color and reset
	diceFaces := []string{"⚀", "⚁", "⚂", "⚃", "⚄", "⚅"}
	return fmt.Sprintf(colorStr, die.Color, diceFaces[roll-1])
}

func main() {
	fmt.Println("Welcome to Terminal Zombie Dice!")
	fmt.Println(`INFO: Type "stop" to end your turn and score your brains.`)
	fmt.Println()
	fmt.Println("Please enter the number of players:")

	var numPlayers int
	_, err := fmt.Scanf("%d\n", &numPlayers)
	if err != nil {
		fmt.Println("Invalid number entered. Exiting.")
		os.Exit(1)
	}

	startGame(numPlayers)
}

func startGame(numPlayers int) {
	players := make([]Player, numPlayers)
	for i := range players {
		players[i].ID = i + 1
	}
	highestScore := 13
	won := false

	// Game loop
	for {

		// Player loop
		for player := range players {

			fmt.Printf("\nPlayer %v's turn!\n", players[player].ID)
			diceBag := newBag()
			drawnDice := [3]Die{}
			brains := Brains{
				Points: 0,
				Dice:   []Die{},
			}
			shotguns := 0
			var answer string
			fmt.Println("Press enter to start rolling:")
			fmt.Scanln()

			// Dice roll loop
			for {
				drawRandomDice(&drawnDice, &diceBag)
				manageDice(&drawnDice, &brains, &shotguns)
				fmt.Printf("Brains: %v, Shotguns: %v\n", brains.Points, shotguns)
				if shotguns >= 3 {
					fmt.Printf("\nOh no! Player %v got shotgunned!\n", players[player].ID)
					brains, shotguns = goNext(players[player].ID, players[player].Score)
					break
				}
				fmt.Println("Roll again or type 'stop' to end your turn:")
				fmt.Scanln(&answer)
				if answer == "stop" {
					players[player].Score += brains.Points
					brains, shotguns = goNext(players[player].ID, players[player].Score)
					break
				}
				if len(diceBag) < 3 {
					diceBag = append(diceBag, brains.Dice...)
					brains.Dice = []Die{}
				}
			}
			if players[player].Score >= highestScore {
				won = true
			}
		}

		if won {
			highestScore, players = getHighestScore(players)
			if len(players) == 1 {
				fmt.Printf("\nPlayer %v wins with %v brains!\n", players[0].ID, players[0].Score)
				break
			}
		}
	}
}

func getHighestScore(players []Player) (int, []Player) {
	highestScore := 0
	var highestPlayers []Player
	for _, player := range players {
		if player.Score > highestScore {
			highestScore = player.Score
			highestPlayers = []Player{player}
		} else if player.Score == highestScore {
			highestPlayers = append(highestPlayers, player)
		}
	}
	return highestScore, highestPlayers
}

func goNext(player, score int) (Brains, int) {
	fmt.Printf("Player %v now has %v brains.\n", player, score)
	return Brains{
		Points: 0,
		Dice:   []Die{},
	}, 0
}

func newBag() []Die {
	const (
		greenDice  = 6
		yellowDice = 4
		redDice    = 3
	)

	diceBag := make([]Die, 0, greenDice+yellowDice+redDice)

	for i := 0; i < greenDice; i++ {
		diceBag = append(diceBag, greenDie)
	}

	for i := 0; i < yellowDice; i++ {
		diceBag = append(diceBag, yellowDie)
	}

	for i := 0; i < redDice; i++ {
		diceBag = append(diceBag, redDie)
	}

	return diceBag
}

func drawRandomDice(drawnDice *[3]Die, bag *[]Die) {
	for i := 0; i < 3; i++ {
		if (*drawnDice)[i] != emptyDie {
			continue
		}
		idx := rand.IntN(len(*bag))
		(*drawnDice)[i] = (*bag)[idx]
		(*bag)[idx] = (*bag)[len(*bag)-1]
		*bag = (*bag)[:len(*bag)-1]
	}
}

func rollDice(die Die) (int, string) {
	face := rand.IntN(6) + 1
	if face <= die.Lose {
		return face, Shotgun
	}
	if face <= die.Win {
		return face, Footprint
	}
	return face, Brain
}

func manageDice(drawnDice *[3]Die, brains *Brains, shotguns *int) {
	for i, die := range *drawnDice {
		roll, result := rollDice(die)
		fmt.Printf("%v = %v\n", die.GetColoredDie(roll), result)

		switch result {
		case Brain:
			brains.Points++
			brains.Dice = append(brains.Dice, die)
			(*drawnDice)[i] = emptyDie
		case Shotgun:
			(*drawnDice)[i] = emptyDie
			(*shotguns)++
		}
	}
}

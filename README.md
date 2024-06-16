# Terminal Zombie Dice

Terminal Zombie Dice is a text-based adaptation of the popular board game "Zombie Dice". It simulates the dice-rolling, risk-taking gameplay of the original game where players take on the role of zombies hungry for brains.

## About the Game

### Gameplay Overview

- **Objective:** The goal is to score the most brains by rolling dice.
- **Players:** Two or more players take turns to roll dice and accumulate points.
- **Dice Types:** There are three types of dice: green (easiest), yellow (medium tough), and red (toughest).
- **Rolling:** Each turn, players shake the cup (or press enter in this terminal version) to draw and roll three dice.
- **Symbols:** Dice have three symbols:
  - **Brain:** You ate your victim's brain.
  - **Shotgun:** The victim fought back!
  - **Footprints:** The victim escaped.
- **Turn Options:** After each roll:
  - If you roll three shotguns, your turn ends with no points scored.
  - Otherwise, decide to stop and score the brains you've collected, or continue rolling for more points.
  - Footprints dice are kept for re-rolling along with new dice.
- **Winning:** The first player to reach or exceed 13 brains triggers the final round. The player with the most brains at the end of that round wins. In case of a tie, only the tied players play a tiebreaker round.

### How to Play

1. **Setup:** Enter the number of players.
2. **Game Rounds:** Players take turns to roll dice, accumulating brains and avoiding three shotguns in a turn.
3. **Scoring:** Each brain rolled scores one point. Decide to stop and score or risk losing points by continuing to roll.
4. **Final Round:** The game ends when a player reaches 13 brains. Finish the round, and the player with the highest score wins.

### Technologies Used

- Written in Go (Golang)
- Uses random number generation for dice rolls
- Simple text-based interface for player interaction

## Acknowledgements

- **Zombie Dice:** Original board game concept by Steve Jackson Games. Visit [zombiedice.sjgames.com](http://zombiedice.sjgames.com) for more information on the official game.
  
## Disclaimer

This project is a fan-made adaptation of the board game "Zombie Dice" by Steve Jackson Games. I am not affiliated with Steve Jackson Games or the official Zombie Dice product. All rights for the original "Zombie Dice" game belong to Steve Jackson Games.

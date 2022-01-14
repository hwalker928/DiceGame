package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

var (
	currentTurn int = 2
	dice1 int
	dice2 int
	p1Name string
	p2Name string
	lastDouble1 bool = false
	lastDouble2 bool = false
	p1Points int = 0
	p2Points int = 0
)

func rollDice() int {
  return rand.Intn(6-1) + 1
}

func clearScreen() {
  screen.Clear()
	screen.MoveTopLeft()
}

func main() {
  // Introduction
  clearScreen()
  color.Magenta("Welcome to the dice game.")
  color.Red("\n\nPlayer 1")
  fmt.Println("It seems we haven't met before.\nWhat is your name?")
  fmt.Scanln(&p1Name) // Take player 1's name and store it as p1Name
  color.Blue("\n\nPlayer 2")
  fmt.Println("It seems we haven't met before.\nWhat is your name?")
  fmt.Scanln(&p2Name) // Take player 2's name and store it as p2Name
  clearScreen()

  for i:=0; i < 32; i++ {
    if currentTurn == 1 {
      currentTurn = 2
    } else {
    currentTurn = 1
    }
    game()
  }

  clearScreen()
  if p1Points > p2Points {
    // Player 1 wins
    color.Red(p1Name, "has won the game with", p1Points, "points!")
  } else if p1Points == p2Points {
    // Ends in a draw
    color.Magenta("The game ended in a draw. Better luck next time.")
  } else {
    // Player 2 wins
    color.Blue(p2Name, "has won the game with", p2Points, "points!")
  }
  fmt.Scanln()
}

func game() {
  clearScreen()
  if currentTurn == 1 {
		str := p1Name + " - " + strconv.Itoa(p1Points)
    color.Red(str)
  } else {
		str := p2Name + " - " + strconv.Itoa(p2Points)
    color.Blue(str)
  }
  color.Magenta("Press ENTER to roll...")
  fmt.Scanln()
  rand.Seed(time.Now().UnixNano()) // Set random seed
  dice1 = rollDice()
  fmt.Println("Dice 1: ", dice1)
  dice2 = rollDice()
  fmt.Println("Dice 2: ", dice2)
  if dice1 == dice2 {
    color.Yellow("\n\n You rolled a double!")
    if currentTurn == 1 {
      p1Points = p1Points + 1
      if lastDouble1 {
        color.Yellow(" 2x in a row! DOUBLE POINTS!")
        p1Points = p1Points * 2
        lastDouble1 = false // If 2x in a row, set last to false to prevent 3x in a row
      } else {
        lastDouble1 = true
      }
    } else {
    p2Points = p2Points + 1
      if lastDouble2 {
        color.Yellow(" 2x in a row! DOUBLE POINTS!")
        p2Points = p2Points * 2
        lastDouble2 = false // If 2x in a row, set last to false to prevent 3x in a row
      } else {
        lastDouble2 = true
      }
    }
  } else {
    if currentTurn == 1 {
      lastDouble1 = false
    } else {
      lastDouble2 = false
    }
  }
  color.Magenta("Press ENTER to end your turn...")
  fmt.Scanln()
}

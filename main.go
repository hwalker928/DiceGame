package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func rollDice() int {
  return rand.Intn(6-1) + 1
}

func clearScreen() {
  screen.Clear()
	screen.MoveTopLeft()
}

func main() {
  // Initiate variables
  var dice1 int
  var dice2 int
  var p1Name string
  var p2Name string
  var lastDouble1 bool = false
  var lastDouble2 bool = false
  var p1Points int = 0
  var p2Points int

  // Add support for colours
  magenta := color.New(color.FgMagenta) // System output
  yellow := color.New(color.FgYellow) // Double output

  red := color.New(color.FgRed) // Player 1 output
  blue := color.New(color.FgCyan) // Player 2 output

	clearScreen()
  magenta.Println("Welcome to the dice game.")
  red.Println("\n\nPlayer 1")
  fmt.Println("It seems we haven't met before.\nWhat is your name?")
  // Take player 1's name and store it as p1Name
  fmt.Scanln(&p1Name)
  blue.Println("\n\nPlayer 2")
  fmt.Println("It seems we haven't met before.\nWhat is your name?")
	// Take player 2's name and store it as p2Name
  fmt.Scanln(&p2Name)
  clearScreen()

  for i:=0; i < 16; i++ {
    clearScreen()
    red.Println(p1Name, "-", p1Points)
    magenta.Println("Press ENTER to roll...")
    fmt.Scanln()
    rand.Seed(time.Now().UnixNano())
    dice1 = rollDice()
    fmt.Println("🎲 Dice 1: ", dice1)
    dice2 = rollDice()
    fmt.Println("🎲 Dice 2: ", dice2)
    if dice1 == dice2 {
      yellow.Println("\n\n You rolled a double! 🎉")
      p1Points = p1Points + 1
      if lastDouble1 {
        yellow.Println(" 2x in a row! DOUBLE POINTS! ✨")
        p1Points = p1Points * 2
        lastDouble1 = false
      } else {
        lastDouble1 = true
      }
    } else {
      lastDouble1 = false
    }
    magenta.Println("Press ENTER to end your turn...")
    fmt.Scanln()

    clearScreen()
    blue.Println(p2Name, "-", p2Points)
    magenta.Println("Press ENTER to roll...")
    fmt.Scanln()
    rand.Seed(time.Now().UnixNano())
    dice1 = rollDice()
    fmt.Println("🎲 Dice 1: ", dice1)
    dice2 = rollDice()
    fmt.Println("🎲 Dice 2: ", dice2)
    if dice1 == dice2 {
      yellow.Println("\n\n You rolled a double! 🎉")
      p2Points = p2Points + 1
      if lastDouble2 {
        yellow.Println(" 2x in a row! DOUBLE POINTS! ✨")
        p2Points = p2Points * 2
        lastDouble2 = false
      } else {
        lastDouble2 = true
      }
    } else {
      lastDouble2 = false
    }
    magenta.Println("Press ENTER to end your turn...")
    fmt.Scanln()
  }

  clearScreen()
  if p1Points > p2Points {
    // Player 1 wins
    red.Println(p1Name, "has won the game with", p1Points, "points!")
  } else if p1Points == p2Points {
    // Ends in a draw
    magenta.Println("The game ended in a draw. Better luck next time.")
  } else {
    // Player 2 wins
    blue.Println(p2Name, "has won the game with", p2Points, "points!")
  }
  fmt.Scanln()
}

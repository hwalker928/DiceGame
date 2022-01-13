package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func rollDice() int {
  return rand.Intn(6-1) + 1
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
  magenta := color.New(color.FgMagenta)
  yellow := color.New(color.FgYellow)

  red := color.New(color.FgRed)
  blue := color.New(color.FgCyan)

  magenta.Println("Welcome to dice game.")
  red.Println("\n\nPlayer 1")
  fmt.Println("It seems we haven't met before.\nWhat is your name?")
  // Take player 1's name and store it as p1Name
  fmt.Scanln(&p1Name)
  blue.Println("\n\nPlayer 2")
  fmt.Println("It seems we haven't met before.\nWhat is your name?")
  fmt.Scanln(&p2Name)
  // It's a mess, I know. Just to clear the screen.
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")
  fmt.Println("\033[2J")

  for i:=0; i < 16; i++ {
    fmt.Println("\033[2J")
    red.Println(p1Name, "-", p1Points)
    magenta.Println("Press ENTER to roll...")
    fmt.Scanln()
    rand.Seed(time.Now().UnixNano())
    dice1 = rollDice()
    fmt.Println("Dice 1: ", dice1)
    dice2 = rollDice()
    fmt.Println("Dice 2: ", dice2)
    if dice1 == dice2 {
      yellow.Println("\n\n You rolled a double!")
      p1Points = p1Points + 1
      if lastDouble1 {
        yellow.Println(" 2x in a row! DOUBLE POINTS!")
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

    fmt.Println("\033[2J")
    blue.Println(p2Name, "-", p2Points)
    magenta.Println("Press ENTER to roll...")
    fmt.Scanln()
    rand.Seed(time.Now().UnixNano())
    dice1 = rollDice()
    fmt.Println("Dice 1: ", dice1)
    dice2 = rollDice()
    fmt.Println("Dice 2: ", dice2)
    if dice1 == dice2 {
      yellow.Println("\n\n You rolled a double!")
      p2Points = p2Points + 1
      if lastDouble2 {
        yellow.Println(" 2x in a row! DOUBLE POINTS!")
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

  fmt.Println("\033[2J")
  if p1Points > p2Points {
    // p1 wins
    red.Println(p1Name, "has won the game with", p1Points, "points!")
  } else if p1Points == p2Points {
    magenta.Println("The game ended in a draw. Better luck next time.")
  } else {
    // p2 wins
    blue.Println(p2Name, "has won the game with", p2Points, "points!")
  }

}

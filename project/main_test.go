package main
import "testing"
import "fmt"
import "./universe"
import "./Exercises"

func success(b bool) {
  fmt.Printf("TECHIO> success %v\n",b)
}

func msg(channel string, msg string) {
  fmt.Printf("TECHIO> message --channel \"%v\" \"%v\"\n", channel, msg)
}

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if observed := HelloWorld(); observed != expected {
    
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
		success(false)
		msg("Oops! 🐞", "Please try again 🤔")
	}else{
		 msg("Kudos 🌟", "You've passed a test!")
   		 success(true)
	}
}

func TestUniverse(t *testing.T) {
  var s = universe.CountAllStars(800,120,450)
  if s != 1370 {
    t.Error("Expected 1370, got ", s)
    success(false)
    msg("Oops! 🐞", "Have you properly accumulated each value into 'total' 🤔")
  } else {
    msg("Kudos 🌟", "You've passed a test!")
    success(true)
  }
}

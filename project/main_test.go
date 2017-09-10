package main
import "testing"
import "fmt"
import "./universe"
import "./exercises"

func success(b bool) {
  fmt.Printf("TECHIO> success %v\n",b)
}

func msg(channel string, msg string) {
  fmt.Printf("TECHIO> message --channel \"%v\" \"%v\"\n", channel, msg)
}

func TestShareWith(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"", "One for you, one for me."},
		{"Alice", "One for Alice, one for me."},
		{"Bob", "One for Bob, one for me."},
	}
	for _, test := range tests {
		if observed := exercises.ShareWith(test.name); observed != test.expected {
			t.Fatalf("ShareWith(%s) = %v, want %v", test.name, observed, test.expected)
			success(false)
			msg("Oops! 👎", "Please try again !")
		}else{
			msg(test.name, test.expected + "👍")
		}
	}
	success(true)
}

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if observed := exercises.HelloWorld(); observed != expected {
    
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

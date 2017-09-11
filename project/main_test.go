package main
import "testing"
import "fmt"
import "./universe"
import "./exercises"
import "strconv"
func success(b bool) {
  fmt.Printf("TECHIO> success %v\n",b)
}

func msg(channel string, msg string) {
  fmt.Printf("TECHIO> message --channel \"%v\" \"%v\"\n", channel, msg)
}


func TestLeapYears(t *testing.T) {
	
	var testCases = []struct {
	year        int
	expected    bool
	description string
	}{
	{2015, false, "year not divisible by 4: common year"},
	{2016, true, "year divisible by 4, not divisible by 100: leap year"},
	{2100, false, "year divisible by 100, not divisible by 400: common year"},
	{2000, true, "year divisible by 400: leap year"},
	}
	
	didPass:=true
	for _, test := range testCases {
		observed := exercises.IsLeapYear(test.year)
		if observed != test.expected {
			t.Fatalf("IsLeapYear(%d) = %t, want %t (%s)",
				test.year, observed, test.expected, test.description)
			  didPass = false
		}else{
			
			msg( fmt.Sprintf("%d%d",test.year,"👍"), strconv.FormatBool(test.expected) )		
			
		}
	}
	
	if !didPass {
	    msg("Oops! 👎", "Please try again !")	
	}
	
	success(didPass)
		
}


func TestShareWith(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"", "One for you, one for me."},
		{"Alice", "One for Alice, one for me."},
		{"Bob", "One for Bob, one for me."},
	}
	didPass:=true
	for _, test := range tests {
		if observed := exercises.ShareWith(test.name); observed != test.expected {
			t.Fatalf("ShareWith(%s) = %v, want %v", test.name, observed, test.expected)
			//success(false)
			didPass = false
			
		}else{
			if test.name == ""{
			      msg("you", test.expected + "👍")	
			}else{
			  msg(test.name, test.expected + "👍")	
			}
			
		}
	}
	
	if !didPass {
	    msg("Oops! 👎", "Please try again !")	
	}
	
	success(didPass)
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

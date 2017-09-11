# Welcome!

This a collection of Exercises to make it easy to practice Go without setting up tools.

If you are new to Go i recommend: Go by Example (https://gobyexample.com), and the official Go tour (https://tour.golang.org) so you can get up and running with the syntax and constructs of the language.



# Hello World
The objective is simple:

Define a function named HelloWorld that takes no arguments,and returns a Hello, World! string.
In other words, define a function with the following signature: HelloWorld() string that returns  Hello, World!


@[The classical introductory exercise. Just say "Hello, World!"]({"stubs":["exercises/Helloworld.go"], "command":"go test -run HelloWorld"})


# Two Fer

Two-fer or 2-fer is short for two for one. One for you and one for me

"One for X, one for me."

When X is a name or "you".

If the given name is "Alice", the result should be "One for Alice, one for me." If no name is given, the result should be "One for you, one for me."

Define a function ShareWith(string) string.

Example:
h := ShareWith("")

fmt.Println(h)

// Output: One for you, one for me.	

@[Two-fer or 2-fer is short for two for one. One for you and one for me]({"stubs":["exercises/twofer.go"], "command":"go test -run ShareWith"})


# Leap Year

Given a year, report if it is a leap year.

The tricky thing here is that a leap year in the Gregorian calendar occurs:

    on every year that is evenly divisible by 4
  

	    except every year that is evenly divisible by 100
  

	    unless the year is also evenly divisible by 400
    
Example:
h := IsLeapYear(2015)

fmt.Println(h)

// Output: false
//reason : "year not divisible by 4: common year"


h := IsLeapYear(2016)

fmt.Println(h)

// Output: true
//reason : ""year divisible by 4, not divisible by 100: leap year""


@[Report if it is a leap year.]({"stubs":["exercises/leapyear.go"], "command":"go test -run LeapYears"})






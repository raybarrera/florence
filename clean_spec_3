:: program main

import {
	"async"
	"foo"
}

type rules trait {
	apply()
}

// Declaring a struct type
type game struct {
	// Declares the `fps` field with an arbitrary tag containing the string "some data"
	// Reflection can pull the metadata in the tag.
	int fps	`some data`
}

// Declaring a type based on game
type fpsGame game {}

// Delcaring a type using generics
type specialGame<rules T> game {
	// T resolves to rules
	T gameRules
}

// call the generic type `gameRules` `apply()` method. 
// This is possible because the specialGame type generic field is constrainted to the rules trait.
specialGame :: fn init() {
	gamerules.apply()
}

type specialRules struct {}

specialRules :: fn apply(){
	print("rules applied.")
}

// Attaching a method, start::(), to a pointer of type game.
// the pointer is available in the scope of the function via keyword "it"
*game :: fn start () {
	// Read as "the val at (@) 'it' (where 'it' is a pointer of type game)"
	@it.fps = 60

	//it.fps = 60 // Compile error. A pointer of type game does not have a field named "fps"
	//The ... operator following a function invocation runs it concurrently
	game.Update()...
}

game :: fn setFps (int newFps) {
	// Since we are operating on a value, there is no need for the @ operator
	it.fps = newFps

	// call method of type game, declared above as a pointer receiver
	it.start()
}

// Implicitly implements the foo::doer trait
game :: fn Update (float64 deltaTime) {
	while true {
		return 
	}
}

fn main (...string argv) {
	c := NewChannel<int32>()
	// Fire off a concurrent invocation of game.Start()
	game.start()...
	
	// Not using concurrency causes this line to block until slowOp() returns
	slowOp(c)
	// The concurrent version below does not block
	slowOp(c)...
	i := <- c //i := 5

	//TODO clear up the blocking/nonblocking nature of channels. The above would not work
}

fn slowOp(chan<int32> ch) {
	sleep(10)
	 5 -> ch
}

fn arrayOfNames() {
	// separate assignment
	var [string] names
	names = ["Ray", "Cara"]

	//or initialize and assign
	lastNames := ["McIlnay", "Barrera"]
}

fn mapOfPeople() {
	var [string, string] people

	persons := [string, string]{
		"some": "one",
	}
}

::package foo

type doer trait {
	Update(float64)
} 
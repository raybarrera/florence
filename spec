// :: scope operator
// program is a keyword indicating this is a program, not a pacakge.
// main is the reserved name for the main activity in a program
:: program main

int count = 1
string name = "abc"
float32 cents

@function
main {
    Person p = {
        age: 20,
        name: "blah",
        lastName: "doo"
    }
    something := p.Count(1, 2)
    other := foo(p)
    log(sum(1,2))
    return 0
}

foo :: (Counter c) int {
    return c.Count(1, 2)
}

// Type declarations
// ..Do we want classes?
// ...Do we want pointers?
// .... Do we want auto-linting like gofmt?
struct Person {
    int     age
    string  name
    string  lastName
}

@struct Person{}

// Contracts, AKA Interfaces (Should I just call them interfaces?) Or traits
@trait Counter {
    Count :: (int, int) int
    Deduct :: (int) int
}

interface Counter {...}

type Counter interface {...}

trait Counter {...}

@trait Counter{...}

@Counter::Count(int, int) int
@Counter::Deduct(int) int


// Methods
Person (ref p) :: Count :: (int a, int b) int {...} 

Person::Count::(int a, int b) int [Person p] {...}

Count::(@Person p, int a, int b) int {...}

// FUNCTION DECLRATIONS -- Operator-inferred function declaration 
sum :: (int a, b) returns int {}

// FUNCTION DECLRATIONS -- Operator-inferred function without "returns" keyword
sum :: (int a, b) int {}

sum :: (int a, b) (int, float32) {}

sum :: (int a, b) int, float32 {}

// FUNCTION DECLARATIONS -- named args?
sum::(int some, int thing;) int{}
sum::(_;int some, int thing) int{}

// Lambdas?
func someFunc := sum :: (int a, b) int{}

someFunc := (int a, b) int{}

// Operator-inferred, using :> instead of ::
foo :> (int a) string {}

// C/go version
func sum(int a, int b) (int) {
    sum := a + b
    return sum
}

func foo() (int a, int b) {
    return 1, 2
}

//  declaring a new scope in the same file.
// This is a scope for package named main_tests
:: package main_tests

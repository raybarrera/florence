// :: scope operator
// program is a keyword indicating this is a program, not a pacakge.
// main is the reserved name for the main activity in a program
:: program main

int count = 1
string name = "abc"
float32 cents

@function
main {
    log(sum(1,2))
    return 0
}

// Annotations version
@function
<- int a, int b
-> int
sum {

}

// Operator-inferred function declaration 
sum :: (int a, b) returns int {

}



// C version
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

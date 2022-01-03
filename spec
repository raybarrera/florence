// :: scope operator
// program is a keyword indicating this is a program, not a pacakge.
// main is the reserved keyword/name for the main activity in a program
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

////////////////////////////////////////////////////
// TYPE DECLARATIONS
////////////////////////////////////////////////////

// ..Do we want classes?
// ...Do we want pointers?
// .... Do we want auto-linting like gofmt?
struct Person {
    int     age
    string  name
    string  lastName
}

@struct Person{}

////////////////////////////////////////////////////
// TRAITS
////////////////////////////////////////////////////

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


////////////////////////////////////////////////////\
// FUNCTIONS
////////////////////////////////////////////////////
Person (ref p) :: Count :: (int a, int b) int {...} 

Person::Count::(int a, int b) int [Person p] {...}

Count::(@Person p, int a, int b) int {...}

// Operator-inferred function declaration 
sum :: (int a, b) returns int {}

// Operator-inferred function without "returns" keyword
sum :: (int a, b) int {}

sum :: (int a, b) (int, float32) {}

sum :: (int a, b) int, float32 {}

// Named args?
sum::(int some, int thing,{int do}) int{}
sum::(int some, int thing) int{}

// C/go version
func sum(int a, int b) (int) {
    sum := a + b
    return sum
}

func foo() (int a, int b) {
    return 1, 2
}
 
////////////////////////////////////////////////////
// POINTERS, VALUES, ETC.
////////////////////////////////////////////////////
// a is an int
sum::(int a) {
    a++ //a is a copy of the passed value, original value remains the same
}

// a is a pointer of type int
//a = 1
//... using implicit casting via default operators
sum::(int@ a) {
    int b := a // 
    b++ //original value is incremented by 1
    a = b // the int at pointer a is now equal to b (2)
}

some::(string@ a) {
    string b := a + "world"
    a = b // OK

}

// ... with dereference operator
some::(string@ a) {
    &a[0] = "a" // err: "string" is immutable
    ...
    string b := &a + " "
    &a = b // OK, new assignment
}

sum::(int@ a) {
    &a++
}

////////////////////////////////////////////////////
// ANONYMOUS FUNCTIONS
////////////////////////////////////////////////////
...
// ... Declaring a function that takes an anonymous function as an argument
someFunc::(int a, func::(int){} f) {
    f(a)
}
someFunc::(int a, func::(int){}foo, func::(int){}bar) {
    foo(a)
    bar(a)
}
someFunc::(int a, func::(int)int{} foo){
    someInt := foo(a)
}

// ... Version without the func keyword
someFunc::(int a, (int){} f) {
    f(a)
}
someFunc::(int a, (int){}foo, (int){}bar) {
    foo(a)
    bar(a)
}
someFunc::(int a, (int)int{} foo){
    i := foo(a)
}

// ... Placeholder version
someFunc::(int a, func f) :: 
(int){} {
    f(a)
}
someFunc::(int a, func foo, func bar)::
    (int){},
    (int){}{
        foo(a)
        bar(a)
    }
someFunc::(int a, func foo)::
    (int)int{}{
        i := foo(a)
    }


//...Inline versions. One and multiple function inputs.

//(string, (int){}){}
someFunc("hello", (1){print()})

//(string, (int){}, (int){}, (int){})
someFunc("hello", 
    (1){print()},
    (2){print()},
    (3){print()})


//... Lambdas that take a placeholder in arguments, and the real anonymous function
// .... in a comma-separated list of blocks after the :: (scope) operator
someFunc("hello", func)::(1){print()}
someFunc("hello", func, func, func)
    ::(1){
        print()
    },
    (1){
        print()
    },
    (1){
        print()
    }


////////////////////////////////////////////////////
// PACKAGE/SCOPE DECLARATION
////////////////////////////////////////////////////
//  declaring a new scope in the same file.
//  NOTES: There is an "open" scope at the top (::), we infer
//  a "closing" scope at end of file, or in this case, a new scope declaration.
//  In theory, this allows a "flat" structure avoiding the unnecessary indentation in languages like C# (prior to v10)
// This is a scope for package named main_tests
:: package main_tests

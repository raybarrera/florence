# Florence Programming Language

Toy language developed using Go's lexer/parser/ast/SSA/compiler as a basis

## Basic program
``` Go
:: Program main

fn main() {
    a := 1 + 2
    print(a) // 3
}
```
## Types
### Struct
``` go
type Foo struct {}
```
#### Attaching a method to a struct
##### By value
```rust
Foo :: fn Do(){}
```

#### Invoking the method
```rust

fn main() {
    foo := Foo{}
    foo.Do()
}
```

##### Using the `it` keyword
```go
Foo :: fn Do() {
    print(it) // "it" refers to the copy of "Foo" passed into this method.
}
```
##### By pointer/ref
```rust
*Bar :: fn Do(){}
```
##### Using the `it` keyword with the `@` dereference operator
```rust
type Bar struct {
    int fps
}

*Bar :: fn Do(){
    // Read as "the val at (@) 'it' (where 'it' is a pointer of type Bar)"
    @it.fps = 60

    ...

    it.fps = 60 // Compile error. A pointer of type Bar does not have a field named "fps"

}
```
 

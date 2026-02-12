# Method Sets

## What It Is

- The **method set** of a type determines which interfaces it satisfies
- **ELI10:** Method sets are the guest list -- a pointer has VIP access to all methods, a value only gets into the value-receiver party.
- `T` has methods with value receivers only
- `*T` has methods with both value and pointer receivers

## Why It Matters

- This determines whether you can assign a value or pointer to an interface variable
- **ELI10:** Get this wrong and the compiler slams the door in your face -- no explanation, just "does not implement."
- One of the trickiest Go concepts and a favorite interview topic

## Syntax Cheat Sheet

```go
type Speaker interface { Speak() string }

type Dog struct { Name string }
func (d Dog) Speak() string { return "Woof" }   // value receiver

type Cat struct { Name string }
func (c *Cat) Speak() string { return "Meow" }  // pointer receiver

var s Speaker
s = Dog{}    // OK: Dog has Speak (value receiver)
s = &Dog{}   // OK: *Dog includes all of Dog's methods
s = &Cat{}   // OK: *Cat has Speak
// s = Cat{} // COMPILE ERROR: Cat does not have Speak (pointer receiver only)
```

**Go vs Python**

```
Go:  var s Speaker = &Cat{}   // *T needed for pointer-receiver methods
Py:  s: Speaker = Cat()       # no distinction, always reference
```

## What main.go Shows

- Value receiver method: both `T` and `*T` satisfy the interface
- Pointer receiver method: only `*T` satisfies the interface
- The compile error you hit when passing a value where a pointer is needed

## Common Interview Traps

- `T` method set: only value-receiver methods
- `*T` method set: value-receiver + pointer-receiver methods
- A value stored in an interface is not addressable -- cannot auto-take `&`
- This is why `Cat{}` (value) cannot satisfy an interface requiring a pointer-receiver method
- Concrete variable auto-addressing works (`cat.Speak()`) but interface assignment does not

## What to Say in Interviews

- "The method set of *T includes all methods, but the method set of T only includes value receivers."
- "A value in an interface is not addressable, so Go cannot auto-take its address."
- "This is why a type with pointer-receiver methods must be assigned to an interface as a pointer."

## Run It

```bash
go run ./03_functions_and_methods/04_method_sets
```

```bash
python ./03_functions_and_methods/04_method_sets/main.py
```

## TL;DR (Interview Summary)

- `T` method set = value-receiver methods only
- `*T` method set = value + pointer receiver methods
- Interface assignment requires the method set to match
- Value in interface is not addressable (no auto `&`)
- Use `*T` (pointer) when assigning to interface with pointer-receiver methods
- Concrete variables auto-address; interface values do not

# Python equivalent of Method Sets (compare with main.go)
# Python has no value/pointer receiver distinction.
# All methods receive a reference (like Go pointer receivers).
# "Duck typing" means any object with the right method satisfies the interface.

from abc import ABC, abstractmethod


class Speaker(ABC):
    @abstractmethod
    def speak(self) -> str:
        ...


class Dog(Speaker):
    def __init__(self, name: str):
        self.name = name

    def speak(self) -> str:
        return f"{self.name} says Woof"


class Cat(Speaker):
    def __init__(self, name: str):
        self.name = name

    def speak(self) -> str:
        return f"{self.name} says Meow"


def make_speak(s: Speaker) -> None:
    print(s.speak())


def main():
    # --- No value/pointer distinction in Python ---
    # Everything is a reference. Both Dog and Cat satisfy Speaker.

    print("--- Dog ---")
    d = Dog("Rex")
    make_speak(d)

    print("\n--- Cat ---")
    c = Cat("Whiskers")
    make_speak(c)

    # In Go, Cat with pointer receiver can only satisfy Speaker as *Cat.
    # In Python, there is no such restriction -- all objects are references.

    # --- Duck typing (Python also supports this without ABC) ---
    print("\n--- Duck typing ---")

    class Duck:
        def speak(self) -> str:
            return "Quack"

    # Duck does not extend Speaker, but has speak() -- duck typing works
    duck = Duck()
    print(duck.speak())

    print("\nKey: Go distinguishes T vs *T method sets. Python does not.")


if __name__ == "__main__":
    main()

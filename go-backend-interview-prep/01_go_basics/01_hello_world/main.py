# Python equivalent of Hello World (compare with main.go)

def main():
    # The simplest Python program.
    # No special entry point required (but we use __main__ convention).
    print("Hello, World!")

    # Print without a trailing newline
    print("Python ", end="")
    print("is ", end="")
    print("interpreted.")

    # Formatted output
    language = "Python"
    year = 1991
    print(f"{language} was released in {year}")


if __name__ == "__main__":
    main()

# Python equivalent of Switch (compare with main.go)
import datetime


def main():
    # --- Python 3.10+ match/case (structural pattern matching) ---
    # Before 3.10: use if/elif chains.
    day = "Wednesday"

    match day:
        case "Monday" | "Tuesday":
            print("early week")
        case "Wednesday":
            print("midweek")
        case "Thursday" | "Friday":
            print("late week")
        case _:
            print("weekend")

    # --- Tagless switch = if/elif chain ---
    hour = datetime.datetime.now().hour

    if hour < 12:
        print("morning")
    elif hour < 17:
        print("afternoon")
    else:
        print("evening")

    # --- Match with variable binding ---
    lang = "Go"
    match lang:
        case "Go":
            print("compiled, statically typed")
        case "Python":
            print("interpreted, dynamically typed")
        case _:
            print("unknown language")

    # --- Python match does NOT fall through (same as Go) ---
    # No need for explicit break or fallthrough.

    # --- Multiple values per case ---
    char = 'e'
    match char:
        case 'a' | 'e' | 'i' | 'o' | 'u':
            print(f"{char} is a vowel")
        case _:
            print(f"{char} is a consonant")


if __name__ == "__main__":
    main()

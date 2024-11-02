from typing import List

def main_kwic() -> None:
    """Main function to handle KWIC processing."""
    lines = read_input()
    shifts = generate_shifts(lines)
    sorted_shifts = alphabetize_shifts(shifts)
    output_shifts(sorted_shifts)

def read_input() -> List[str]:
    """Reads input lines from the user until an empty line is entered."""
    print("Enter lines of text (press Enter on an empty line to stop):")
    lines = []
    while True:
        line = input().strip()
        if line == "":
            break
        lines.append(line)
    return lines

def generate_shifts(lines: List[str]) -> List[str]:
    """Generates all circular shifts for each line."""
    shifts = []
    for line in lines:
        words = line.split()
        for i in range(len(words)):
            shift = " ".join(words[i:] + words[:i])
            shifts.append(shift)
    return shifts

def alphabetize_shifts(shifts: List[str]) -> List[str]:
    """Sorts the shifts alphabetically."""
    return sorted(shifts)

def output_shifts(shifts: List[str]) -> None:
    """Outputs the sorted shifts."""
    print("\nAlphabetized Shifts:")
    for shift in shifts:
        print(shift)

if __name__ == "__main__":
    main_kwic()

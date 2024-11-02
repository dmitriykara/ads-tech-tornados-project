# Task 8 solution by Dmitriy Kara

## Project Structure

```
├── README.md
└── kwic
    ├── kwic.py
    ├── test_kwic.py
└── eight_queens
    ├── eight_queens.py
    ├── test_eight_queens.py
```

## Project Description
This project contains solutions to two classical software design problems using **Method 2: Main/Subroutine with Stepwise Refinement**:
1. **Key Word in Context (KWIC)**
2. **Eight Queens Puzzle**

The chosen method satisfies project requirements, focusing on decomposition into modular subroutines to facilitate iterative development and efficient problem-solving.

## Running the Programs
### KWIC Solution
1. Navigate to the `kwic` folder.
2. Run `kwic.py` using `python kwic.py`.
3. Enter lines of text, and press **Enter** on an empty line to finish input.
4. The program will output circular shifts alphabetically.

### Eight Queens Solution
1. Navigate to the `eight_queens` folder.
2. Run `eight_queens.py` using `python eight_queens.py`.
3. The program will display all valid solutions to the Eight Queens puzzle.

## Running Tests
### KWIC Tests
- Run `test_kwic.py` using `python -m unittest kwic/test_kwic.py`.

### Eight Queens Tests
- Run `test_eight_queens.py` using `python -m unittest eight_queens/test_eight_queens.py`.

## Design Approach and Justification
**Method 2: Main/Subroutine with Stepwise Refinement** was used to meet the project’s requirement to facilitate comparison. This method structures the code into a main function and a set of well-defined subroutines, making the design manageable and efficient for both problems.

---

### Comments on Code, Solution, and Testing

1. **Method Selection**: **Method 2: Main/Subroutine with Stepwise Refinement** was chosen because it allows clear decomposition of tasks and aligns with the requirement to enable a comparative study with other methods.
2. **KWIC Implementation**: The KWIC solution processes input, generates circular shifts, sorts them, and outputs results. The subroutine-based design ensures a clear separation of tasks, making the code easy to follow but somewhat rigid for future extensions.
3. **Eight Queens Implementation**: A recursive backtracking algorithm was used, with subroutines handling queen placement and safety checks. This design is well-suited for the problem, ensuring efficiency and ease of optimization.
4. **Testing**: Unit tests were created for key functions in both problems using Python’s `unittest` framework. This helps ensure code correctness and facilitates future modifications.
5. **Project Documentation**: The README.md file provides clear instructions for running the code and tests.

---

### Comparison Tables

ToDo

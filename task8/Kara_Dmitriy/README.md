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

#### Problem A: KWIC (Key Word in Context)

| **Criteria**                                | **Pipes-and-Filters**        | **Abstract Data Types**     | **Main/Subroutine**        | **Implicit Invocation**   |
|---------------------------------------------|------------------------------|-----------------------------|----------------------------|---------------------------|
| **Ease of changing implementation algorithm** | Moderate: Modules are loosely coupled but require significant refactoring if data flow changes | Moderate: Requires changing method implementations but overall structure remains intact | Low: Algorithm changes need updates in several functions, making it less adaptable | High: Event handlers can be added or modified with minimal effort |
| **Ease of changing data representation**      | High: Easy to add filters or change representation at any stage | Moderate: Representation changes require updating related class methods | Low: Data changes require updates across functions, reducing flexibility | High: Data changes easily handled by updating event listeners |
| **Ease of adding additional functions**       | High: New functions can be added as separate filters | Low: New functions require integrating into the object hierarchy | Low: New features require restructuring existing functions | High: New event handlers can be added easily |
| **Performance**                               | Moderate: Data copying between stages can impact performance | High: Efficient, but performance depends on class design | High: Efficient with a clear flow of execution | Moderate: Potential overhead from event-driven mechanisms |
| **Reusability for similar problems**          | High: Modular design makes it adaptable | Moderate: Somewhat reusable but tied to specific implementations | Low: Designed specifically for this problem, hard to adapt | High: Easily reusable by modifying event triggers and listeners |

---

#### Problem B: Eight Queens

| **Criteria**                                | **Pipes-and-Filters**        | **Abstract Data Types**     | **Main/Subroutine**        | **Implicit Invocation**   |
|---------------------------------------------|------------------------------|-----------------------------|----------------------------|---------------------------|
| **Ease of changing implementation algorithm** | Low: Unsuitable for recursive problems, hard to modify | Moderate: Algorithm updates require structural changes | High: Easy to modify or extend the algorithm, e.g., adding more constraints | Moderate: Changes require updating event handlers, which can be manageable |
| **Ease of changing data representation**      | Low: Poor fit for board representation | Moderate: Changes require updating related methods and attributes | Moderate: Board representation can be changed with minor updates to safety checks | Low: Limited flexibility as event-driven updates are not optimal for this problem |
| **Ease of adding additional functions**       | Low: Difficult to add features like visualizations | Moderate: Possible but requires integrating into class structure | High: Additional features like visualizations can be added easily | Moderate: Additional event handlers possible but may introduce complexity |
| **Performance**                               | Low: Inefficient for solving combinatorial problems | High: Well-optimized, depending on class design | High: Recursive backtracking is efficient | Moderate: Event-driven execution introduces overhead |
| **Reusability for similar problems**          | Low: Not reusable for other algorithmic challenges | Moderate: Reusable with some modifications to class design | High: Backtracking algorithm can be reused for similar puzzles | Moderate: Reusable but not optimal for combinatorial problems |

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
This project contains solutions to two classical software design problems:
1. **Key Word in Context (KWIC)** using **Main/Subroutine**.
2. **Eight Queens Puzzle** using **Abstract Data Types (ADT)**.

## Running the Programs
### KWIC Solution
1. Navigate to the `kwic` folder.
2. Run `kwic.py` using `python kwic.py`.
3. Enter lines of text, and press **Enter** on an empty line to finish input.
4. The program will output circular shifts alphabetically.

### Eight Queens Solution
1. Navigate to the `eight_queens` folder.
2. Run `eight_queens_adt.py` using `python eight_queens_adt.py`.
3. The program will display all valid solutions to the Eight Queens puzzle.

## Running Tests
### KWIC Tests
- Run `test_kwic.py` using `python -m unittest kwic/test_kwic.py`.

### Eight Queens Tests
- Run `test_eight_queens_adt.py` using `python -m unittest eight_queens/test_eight_queens_adt.py`.

## Design Approach and Justification
### KWIC: Main/Subroutine Method
- **Justification**: The Main/Subroutine method was chosen for simplicity and efficiency. It clearly separates the text processing tasks into manageable functions, making the solution easy to follow and debug.
- **Structure**: The KWIC solution consists of functions for reading input, generating circular shifts, sorting them, and outputting the results. The design is straightforward but less flexible for future extensions.

### Eight Queens: Abstract Data Types (ADT)
- **Justification**: Using ADT allows encapsulation of the board and queen logic, promoting better data management and a clean object-oriented approach. The QueenBoard class manages the state and provides methods for checking safety, placing, and removing queens.
- **Structure**: The Eight Queens solution uses a recursive backtracking algorithm, leveraging the QueenBoard class to simplify board operations and ensure code clarity.

### Comments on Code, Solution, and Testing
- **Code**: Both solutions are well-structured, with clear separation of concerns. The KWIC code uses procedural programming, while the Eight Queens code adopts an object-oriented design.
- **Solution**: The KWIC solution efficiently generates and sorts shifts, while the Eight Queens solution leverages recursion and ADT for optimal performance.
- **Testing**: Unit tests are provided for key functionalities, ensuring correctness and making the solutions easier to maintain and extend.

### Comparison Tables

#### Problem A: KWIC (Key Word in Context)

| **Criteria**                                | **Pipes-and-Filters**        | **Abstract Data Types**     | **Main/Subroutine**        | **Implicit Invocation**   |
|---------------------------------------------|------------------------------|-----------------------------|----------------------------|---------------------------|
| **Ease of changing implementation algorithm** | Moderate: Modules are loosely coupled but require significant refactoring if data flow changes | Moderate: Requires changing method implementations but overall structure remains intact | Low: Algorithm changes need updates in several functions, making it less adaptable | High: Event handlers can be added or modified with minimal effort |
| **Ease of changing data representation**      | High: Easy to add filters or change representation at any stage | Moderate: Representation changes require updating related class methods | Low: Data changes require updates across functions, reducing flexibility | High: Data changes easily handled by updating event listeners |
| **Ease of adding additional functions**       | High: New functions can be added as separate filters | Low: New functions require integrating into the object hierarchy | Low: New features require restructuring existing functions | High: New event handlers can be added easily |
| **Performance**                               | Moderate: Data copying between stages can impact performance | High: Efficient, but performance depends on class design | High: Efficient with a clear flow of execution | Moderate: Potential overhead from event-driven mechanisms |
| **Reusability for similar problems**          | High: Modular design makes it adaptable | Moderate: Somewhat reusable but tied to specific implementations | Low: Designed specifically for this problem, hard to adapt | High: Easily reusable by modifying event triggers and listeners |

#### Problem B: Eight Queens

| **Criteria**                                | **Pipes-and-Filters**        | **Abstract Data Types**     | **Main/Subroutine**        | **Implicit Invocation**   |
|---------------------------------------------|------------------------------|-----------------------------|----------------------------|---------------------------|
| **Ease of changing implementation algorithm** | Low: Unsuitable for recursive problems, hard to modify | High: Easy to modify or extend board operations | High: Straightforward to update or extend, ideal for backtracking | Moderate: Requires updating event handlers, which can be manageable |
| **Ease of changing data representation**      | Low: Poor fit for board representation | High: Encapsulated in classes, easy to modify | Moderate: Board representation changes require more code updates | Low: Limited flexibility, event-driven mechanisms aren't ideal for this problem |
| **Ease of adding additional functions**       | Low: Difficult to add features like visualizations | High: New methods can be easily added to the class | High: Straightforward to add features like printing solutions | Moderate: Possible but may add complexity |
| **Performance**                               | Low: Inefficient for solving combinatorial problems | High: Optimized with object-oriented design | High: Efficient recursive backtracking | Moderate: Event-driven execution introduces some overhead |
| **Reusability for similar problems**          | Low: Not reusable for other algorithmic challenges | High: Encapsulation makes it adaptable | High: Well-suited for similar recursive problems | Moderate: Limited reusability in a non-event-driven context |

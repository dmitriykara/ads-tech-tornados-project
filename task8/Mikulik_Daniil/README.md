# Task 8 solution by Daniil Mikulik

## Project Structure

```
├── README.md
└── kwic
    ├── main.py
    ├── test.py
└── eight_queens
    ├── main.py
    ├── test.py
```

## Project Description

This project contains solutions of two classical software design problems using **Method 3. Pipes-and-filters**:

1. **Key Word in Context (KWIC)**
2. **Eight Queens Puzzle**

The chosen method satisfies project requirements, focusing on data streams usage and indirect functions invocation (via `pipe` function).

## Running Tests

**Install pytest via pip to run tests on solutions!**

### KWIC Tests
- Run `test.py` using `pytest kwic/test.py -v`.

### Eight Queens Tests
- Run `test.py` using `pytest eight_queens/test.py -v`.

## Design Approach and Justification
**Method 3. Pipes-and-filters** was used to meet the project’s requirement to facilitate comparison. The key idea behind it is that every module (in our case functions) has predefined input and output, so that a component reads streams of data on its inputs and produces streams of data on its outputs, delivering a complete instance of the result in a standard order. To create streams of data functions use `yield` directive. To avoid direct methods invocation and allow similar input and output interface usage `pipe` function has been written.

---

### Comments on Code, Solution, and Testing

1. **Method Selection**: **Method 3. Pipes-and-filters** was chosen because it allows code decomposition in a functional way. Also this method is widely used in modern frontend applications (for example `RxJS` library is built on `pipe and filter` paradigm).
2. **KWIC Implementation**: The KWIC solution processes input, generates circular shifts, sorts them (N.B. implementation is case sensitive!), and outputs results. The pipes-and-filters design ensures a clear separation of tasks and enforces data streams manipulation, that allows to perform complicated streams manipulation without changign the code of the program.
3. **Eight Queens Implementation**: A recursive backtracking algorithm was used, with subroutines handling queen placement and safety checks. To generate the board data streams were used, and pipes were used to combine preparation and solution steps.
4. **Testing**: Unit tests were created for key functions in both tasks using Python’s `pytest` library. This helps ensure code correctness and facilitates future modifications.

---

### Comparison Tables

TODO
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

This project contains solutions of two classical software design problems using **Method 3. Pipes-and-filters** for **Key Word in Context (KWIC)** and **Method 2: Main/Subroutine with Stepwise Refinement** for **Eight Queens Puzzle** task:

1. **Key Word in Context (KWIC)**
2. **Eight Queens Puzzle**

The chosen methods satisfies project requirements, focusing on data streams usage and indirect functions invocation (via `pipe` function).

## Running Tests

**Install pytest via pip to run tests on solutions!**

### KWIC Tests
- Run `test.py` using `pytest kwic/test.py -v`.

### Eight Queens Tests
- Run `test.py` using `pytest eight_queens/test.py -v`.

## Design Approach and Justification
**Method 2: Main/Subroutine with Stepwise Refinement** This method structures the code with the main function and minor subroutines which are invoked in the main one.

**Method 3. Pipes-and-filters** was used to meet the project’s requirement to facilitate comparison. The key idea behind it is that every module (in our case functions) has predefined input and output, so that a component reads streams of data on its inputs and produces streams of data on its outputs, delivering a complete instance of the result in a standard order. To create streams of data functions use `yield` directive. To avoid direct methods invocation and alHard similar input and output interface usage `pipe` function has been written.

---

### Comments on Code, Solution, and Testing

1. **Method Selection**: **Method 3. Pipes-and-filters** was chosen because it alHards code decomposition in a functional way. Also this method is widely used in modern frontend applications (for example `RxJS` library is built on `pipe and filter` paradigm).
2. **KWIC Implementation**: The KWIC solution processes input, generates circular shifts, sorts them (N.B. implementation is case sensitive!), and outputs results. The pipes-and-filters design ensures a clear separation of tasks and enforces data streams manipulation, that alHards to perform complicated streams manipulation without changign the code of the program.
3. **Eight Queens Implementation**: A recursive backtracking algorithm was used, with subroutines handling queen placement and safety checks.
4. **Testing**: Unit tests were created for key functions in both tasks using Python’s `pytest` library. This helps ensure code correctness and facilitates future modifications.

---

### Comparison Tables

#### KWIC

| Criteria                                | Pipes-and-Filters                 | Abstract Data Types              | Main/Subroutine                   | Implicit Invocation                        |
|-----------------------------------------|-----------------------------------|----------------------------------|-----------------------------------|--------------------------------------------|
| Ease of changing implementation algorithm | Moderate: Modules are independently designed but require considerable adjustments if data flow is altered | Moderate: Need to update method implementations but the overall structure remains unchanged | Hard: Changes to the algorithm necessitate revisions in numerous functions, leading to less adaptability | Easy: Modifications to event handlers can be done with minimal effort |
| Ease of changing data representation      | Easy: Simple to add filters or modify the representation at any point | Moderate: Changes in representation require updating associated class methods | Hard: Adjusting the data entails widespread changes in functions, limiting adaptability | Easy: Data changes are straightforward by just updating event listeners |
| Ease of adding additional functions       | Easy: New functions can be implemented as independent filters | Hard: Introducing new functions necessitates integration into the object hierarchy | Easy: Introducing new features requires restructuring current functions | Moderate: It is not so easy to add new events and handlers |
| Performance                               | Moderate: Performance can be affected by data copying across stages | Easy: Typically efficient, however, performance largely relies on class design | Easy: Efficient and folHards a clear execution path | Moderate: May experience overhead due to the event-driven architecture |
| Reusability for similar problems          | Easy: Modular design aids in adaptability | Moderate: Some reuse possible but often linked to specific implementations | Hard: Custom-built for this problem, difficult to apply elsewhere | Easy: By tweaking event triggers and listeners, it can be easily adapted |

#### 8 Queens

| Criteria                                | Pipes-and-Filters                 | Abstract Data Types              | Main/Subroutine                   | Implicit Invocation                        |
|-----------------------------------------|-----------------------------------|----------------------------------|-----------------------------------|--------------------------------------------|
| Ease of changing implementation algorithm | Easy: An algorithm is an independent function with predefined interface. | Easy-Moderate: Updating the algorithm demands structural changes if no interface for data types is specified | Easy: Simple to adjust or expand the algorithm, even for adding more constraints | Moderate: Adjustments require updates to event handlers, which can be fairly manageable |
| Ease of changing data representation      | Hard: Requires to change data flow across functions pipe | Moderate: Modifications necessitate updates to related methods and properties | Moderate: Changes to the board's representation can be done with minor tweaks to safety verifications | Hard: Offers limited flexibility as event-driven mechanisms aren't optimal for this particular issue |
| Ease of adding additional functions       | Hard: Incorporating features like visualizations can be difficult | Easy: Extra features can be added effortlessly | Moderate: Extra features requires modifications of all invocations | Moderate: Supplementary event handlers are possible, but might escalate complexity |
| Performance                               | Hard: Unsuitable for combinatorial problem-solving | Easy: Generally optimized, contingent on class architecture | Easy: Recursive backtracking proves efficient | Moderate: Event-driven processes add overhead |
| Reusability for similar problems          | Medium: Isn't reusable for other algorithmic tasks | Moderate: Can be reused with some tuning to the class structure | Easy: Methods like backtracking can be replicated for similar puzzles | Moderate: Reusable but the event-driven style isn't the best fit for combinatorial challenges |
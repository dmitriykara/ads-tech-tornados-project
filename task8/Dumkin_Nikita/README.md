# Task 8 solution by Dumkin Nikita

## Project Structure

```
├── README.md
└── kwic
    ├── index.py
└── eight_queens
    ├── index.py
```

## Project Description

This project contains solutions of two classical software design problems using **Method 4. Implicit invocation** and **Method 2. Main/Subroutine with stepwise refinement** respectively:

1. **Key Word in Context (KWIC)**
2. **Eight Queens Puzzle**

## Running

### KWIC run
- Run `index.py` using `python index.py`.

### Eight Queens run
- Run `index.py` using `python index.py`.

## Design Approach and Justification
**Method 4. Implicit invocation**
The `Event` class simulates an event system. Each event holds a list of subscribers (functions that should be called when the event is triggered). When notify() is called, it triggers all subscribers with the given data.

`KWICSystem`: The main KWICSystem class contains: 
- `Events`: Three events (input_received_event, circular_shift_event, and alphabetize_event) to handle different stages of the KWIC process.
- `Subscribers`: Methods like circular_shift and alphabetize subscribe to these events.
- `Steps`: The system uses three main methods — circular_shift, alphabetize, and display_results—to process the text input in stages.

This approach achieves a loosely coupled, event-driven solution to the KWIC problem, where each processing step is independent and triggered via event notifications

---

**Method 2. Main/Subroutine with stepwise refinement** To solve this problem using a Main/Subrtine design with stepwise refinement and shared data, it was necessary to break the solution into smaller, manageable functions:

Main Function: initializes the chessboard, calls the solve function, and prints the result.

Subroutines:
- `solve`: a recursive function that attempts to place queens on the board.
- `is_safe`: checks if it is safe to place a queen in a certain position.
- `print_board`: prints the board with the queens placed on it.
The solution uses a shared variable **board** that will be updated as queens are placed or removed.

---

### Comparison Tables

#### KWIC

|     Aspect     |                          Abstract Data Types (ADT)                          |        Main/Subroutine with Stepwise Refinement (Shared Data)        |                        Pipes-and-Filters                        |              Implicit Invocation (Event-Driven)              |
|:--------------:|:---------------------------------------------------------------------------:|:--------------------------------------------------------------------:|:---------------------------------------------------------------:|:------------------------------------------------------------:|
| Design Pattern | Encapsulation of data and operations                                        | Functional decomposition with central shared data structure          | Independent processing stages connected in a linear pipeline    | Event-driven with independent modules subscribing to events  |
| Modularity     | High modularity, encapsulates functionality and data access                 | Limited, as modules share data and interact through subroutine calls | High, each stage is an independent module                       | High, modules communicate indirectly via event notifications |
| Coupling       | Low coupling between data and operations, but tight within an ADT structure | Tightly coupled due to shared data structures and sequential calls   | Low coupling; each stage only needs to know input/output format | Loose coupling; modules don’t know other modules directly    |
| Ease of Modification | Easy to modify within ADT boundaries, but may need to add new ADTs | Difficult; modifications in one subroutine may affect others due to shared data | Easy to modify stages independently                          | Easy; new modules can subscribe to events or existing modules can be updated |

#### 8 Queens

|     Aspect     |                              Abstract Data Types (ADT)                              |                    Main/Subroutine with Stepwise Refinement (Shared Data)                    |                              Pipes-and-Filters                             |                            Implicit Invocation (Event-Driven)                           |
|:--------------:|:-----------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------:|
| Design Pattern | Encapsulates board and queen placement logic into reusable classes                  | Sequential subroutines to place queens, check validity, and backtrack with shared board data | Filters for board configuration, placement, validity, and solution display | Event-driven with independent modules handling each stage (e.g., placement, validation) |
| Modularity     | High; board state and queen placement logic are encapsulated in ADTs                | Limited; all operations are tightly coupled through shared board data                        | Moderate; each stage encapsulates a part of the process                    | High; each stage triggers events without knowing other stages                           |
| Coupling       | Low within each ADT; components are tightly coupled to board and validation methods | High, as subroutines share data and must execute in sequence                                 | Low; only need input/output compatibility between filters                  | Low; modules are only loosely connected through events                                  |
| Ease of Modification | Easy to modify within ADT structure, with minimal impact on other parts | Difficult; changes in one subroutine can affect others due to shared data                    | Easy; each filter can be modified independently                            | Easy; modules can be added or removed as needed                                         |

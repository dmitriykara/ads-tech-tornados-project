## Launch
Download the file Problems_A,_B.ipynb, then go to Google Colab https://colab.google/ and upload the file.

## Design Approach and Justification

### Key Word in Context (KWIC)

### Method 1: Abstract Data Types

The Abstract Data Types (ADT) method uses object-oriented principles to encapsulate functionality and data for the KWIC problem. Key classes include:

- **TextProcessor**:
  - Cleans and prepares text for indexing by removing stop words and splitting text into words.

- **Context**:
  - Stores the context of keywords, including surrounding words and page number, and formats output for display.

- **KWICIndex**:
  - Manages the KWIC index by storing keywords and their contexts. It generates index entries by processing text through the TextProcessor.

- **KWICDisplay**:
  - Handles the display of the KWIC index in a user-friendly format.

The design promotes encapsulation, reusability, and separation of concerns, making it easier to maintain and adapt the system.

---

### Eight Queens (8Q)

#### Method 3: Pipes-and-Filters

The Pipes-and-Filters method for the Eight Queens problem structures the solution as a series of independent processing stages:

- **NQueens Class**:
  - Implements the N Queens solver, handling initialization, queen placement, and solution validation through a recursive method.

### Methods:
- `initialize`: Sets up the board.
- `place_queens`: Recursively places queens, storing valid configurations.
- `is_safe`: Checks if a queen can be placed in a specified position.
- `display_solutions`: Formats and prints all found solutions.

This architecture allows for clear data flow through distinct steps, enhancing modularity and flexibility, with the potential for parallel processing and easy integration of new filters if needed.




## Comparison with the work of Dmitry Kara

### Problem A. Key Word in Context (KWIC)
| **Criteria**                           | **Pipes-and-Filters (KWIC)**                         | **Abstract Data Types  (KWIC)**                                   |
|----------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------------|
| **Ease of changing implementation**    | Moderate: Functions (or "filters") are modular, allowing changes to individual steps, but changes to data flow require multiple adjustments. | Moderate: Requires modifying class methods, but overall structure remains intact. | High: New event handlers can be added or modified with minimal effort.  |
| **Ease of changing data representation** | High: Filters can be adjusted or replaced to handle different data representations without impacting the entire flow. | Moderate: Changes in data representation require updates to class methods but retain overall structure. |
| **Ease of adding additional functions** | High: New processing functions can be added as separate filters without modifying existing ones. | Low to Moderate: New functionality requires integration within the existing class structure, potentially affecting multiple components. |
| **Performance**                        | Moderate: Data passes sequentially through filters, which may introduce redundancy and extra processing time on large inputs. | Moderate: Efficiency is dependent on class design and method interactions. |
| **Reusability for similar problems**   | High: Modular filter-based design is highly adaptable to other pipeline-based text processing tasks. | Moderate: Class-based structure is reusable but may need substantial adaptation for different applications. |

### Problem B. Eight Queens (8Q)
| Criteria                           | Pipes-and-Filters (8Q)                             | Main/Subroutine with Stepwise Refinement  (8Q)      |
|------------------------------------|------------------------------------------------|-------------------------------------------------|
| **Ease of changing implementation algorithm** | Moderate: Modules are loosely coupled, but changes in data flow require significant refactoring. | Moderate: Requires modifying class methods, but overall structure remains intact |
| **Ease of changing data representation** | High: Filters can be modified or swapped to change data representation with minimal impact on other components. | Low: Data changes need to be accounted for across multiple functions, reducing flexibility. |
| **Ease of adding additional functions** | High: New functions can be added as independent filters without impacting existing ones. | Low: Adding new features requires integrating them into the main flow, often restructuring existing functions. |
| **Performance**                   | Moderate: Data copying between stages can impact performance, though parallelization is possible. | High: Direct control of function calls leads to efficient execution with minimal overhead. |
| **Reusability for similar problems** | High: Modular and filter-based design makes the approach highly adaptable to similar problems. | Low: Code is closely tailored to this specific problem, making it harder to adapt for other uses. |


## Comparison with the work of Dumkin Nikita


### Problem A. Key Word in Context (KWIC)
| Criteria                           | Abstract Data Types     (KWIC)                         | Implicit Invocation    (KWIC)                        |
|------------------------------------|-------------------------------------------------|------------------------------------------------|
| **Ease of changing implementation algorithm** | Moderate: Requires modifying class methods, but overall structure remains intact. | High: New event handlers can be added or modified with minimal effort. |
| **Ease of changing data representation** | Moderate: Changes in data representation require updates to class methods but retain overall structure. | High: Changes can be easily managed through event notifications without affecting core logic. |
| **Ease of adding additional functions** | Low: Adding new features often requires integrating them into the existing class hierarchy. | High: New functionalities can be added as new events or subscribers without altering existing code. |
| **Performance**                   | Moderate: Efficiency is dependent on class design and method interactions. | Moderate: Event-driven mechanisms may introduce some overhead but can optimize performance through decoupled components. |
| **Reusability for similar problems** | Moderate: Class-based structure is reusable but may need substantial adaptation for different applications. | High: The modular design facilitates reuse across different contexts by simply modifying event triggers and handlers. |


### Problem B. Eight Queens (8Q)
| Criteria                           | Pipes-and-Filters (8Q)                          | Main/Subroutine with Stepwise Refinement (8Q) |
|------------------------------------|-------------------------------------------------|-------------------------------------------------|
| **Ease of changing implementation algorithm** | Moderate: Modules are loosely coupled, but changes in data flow require significant refactoring. | Low: Algorithm changes require updates across several functions, making adaptation complex. |
| **Ease of changing data representation** | High: Filters can be modified or swapped to change data representation with minimal impact on other components. | Low: Data changes need to be accounted for across multiple functions, reducing flexibility. |
| **Ease of adding additional functions** | High: New functions can be added as independent filters without impacting existing ones. | Low: Adding new features requires integrating them into the main flow, often restructuring existing functions. |
| **Performance**                   | Moderate: Data copying between stages can impact performance, though parallelization is possible. | High: Direct control of function calls leads to efficient execution with minimal overhead. |
| **Reusability for similar problems** | High: Modular and filter-based design makes the approach highly adaptable to similar problems. | Low: Code is closely tailored to this specific problem, making it harder to adapt for other uses. |


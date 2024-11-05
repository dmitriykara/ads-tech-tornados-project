## Launch
Download the file Problems_A,_B.ipynb, then go to Google Colab https://colab.google/ and upload the file.

## Comparison with the work of Dmitry Kara
| **Criteria**                           | **Pipes-and-Filters (KWIC)**                         | **Abstract Data Types Solution (KWIC)**                                   |
|----------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------------|
| **Ease of changing implementation**    | Moderate: Functions (or "filters") are modular, allowing changes to individual steps, but changes to data flow require multiple adjustments. | High: The class-based design with specific methods allows localized changes, making it easier to adapt individual components. |
| **Ease of changing data representation** | High: Filters can be adjusted or replaced to handle different data representations without impacting the entire flow. | Moderate: The data representation is embedded in class attributes, so changes require updates in multiple methods, but classes contain the impact. |
| **Ease of adding additional functions** | High: New processing functions can be added as separate filters without modifying existing ones. | Low to Moderate: New functionality requires integration within the existing class structure, potentially affecting multiple components. |
| **Performance**                        | Moderate: Data passes sequentially through filters, which may introduce redundancy and extra processing time on large inputs. | High: Encapsulated class methods improve efficiency and reduce redundancy, supporting better scalability. |
| **Reusability for similar problems**   | High: Modular filter-based design is highly adaptable to other pipeline-based text processing tasks. | Moderate: Class-based structure is reusable but may need substantial adaptation for different applications. |


| Criteria                           | Pipes-and-Filters (8Q)                             | Main/Subroutine with Stepwise Refinement  (8Q)      |
|------------------------------------|------------------------------------------------|-------------------------------------------------|
| **Ease of changing implementation algorithm** | Moderate: Modules are loosely coupled, but changes in data flow require significant refactoring. | Low: Algorithm changes require updates across several functions, making adaptation complex. |
| **Ease of changing data representation** | High: Filters can be modified or swapped to change data representation with minimal impact on other components. | Low: Data changes need to be accounted for across multiple functions, reducing flexibility. |
| **Ease of adding additional functions** | High: New functions can be added as independent filters without impacting existing ones. | Low: Adding new features requires integrating them into the main flow, often restructuring existing functions. |
| **Performance**                   | Moderate: Data copying between stages can impact performance, though parallelization is possible. | High: Direct control of function calls leads to efficient execution with minimal overhead. |
| **Reusability for similar problems** | High: Modular and filter-based design makes the approach highly adaptable to similar problems. | Low: Code is closely tailored to this specific problem, making it harder to adapt for other uses. |


## Comparison with the work of Mikulik Daniil


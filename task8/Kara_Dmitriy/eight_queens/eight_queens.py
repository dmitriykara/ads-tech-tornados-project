from typing import List

def main_eight_queens() -> None:
    """Main function to initiate solving the Eight Queens problem."""
    board = [-1] * 8
    solve_queens(board, 0)

def solve_queens(board: List[int], row: int) -> None:
    """Recursive backtracking function to place queens."""
    if row == 8:
        print_board(board)
        return

    for col in range(8):
        if is_safe(board, row, col):
            board[row] = col
            solve_queens(board, row + 1)
            board[row] = -1  # Backtrack

def is_safe(board: List[int], row: int, col: int) -> bool:
    """Checks if it's safe to place a queen at (row, col)."""
    for i in range(row):
        if board[i] == col or \
           board[i] - i == col - row or \
           board[i] + i == col + row:
            return False
    return True

def print_board(board: List[int]) -> None:
    """Outputs the board configuration."""
    for row in range(8):
        line = ""
        for col in range(8):
            line += "Q " if board[row] == col else ". "
        print(line)
    print("\n")

if __name__ == "__main__":
    main_eight_queens()

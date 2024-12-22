from typing import List

class QueenBoard:
    def __init__(self, size: int = 8):
        self.size = size
        self.board = [-1] * size

    def is_safe(self, row: int, col: int) -> bool:
        """Check if it's safe to place a queen at (row, col)."""
        for i in range(row):
            if self.board[i] == col or \
               self.board[i] - i == col - row or \
               self.board[i] + i == col + row:
                return False
        return True

    def place_queen(self, row: int, col: int) -> None:
        """Place a queen at (row, col)."""
        self.board[row] = col

    def remove_queen(self, row: int) -> None:
        """Remove the queen from (row, col)."""
        self.board[row] = -1

    def print_board(self) -> None:
        """Print the board configuration."""
        for row in range(self.size):
            line = ""
            for col in range(self.size):
                line += "Q " if self.board[row] == col else ". "
            print(line)
        print("\n")

def solve_queens(board: QueenBoard, row: int) -> None:
    """Recursive backtracking function to place queens."""
    if row == board.size:
        board.print_board()
        return

    for col in range(board.size):
        if board.is_safe(row, col):
            board.place_queen(row, col)
            solve_queens(board, row + 1)
            board.remove_queen(row)

def main_eight_queens() -> None:
    """Main function to initiate solving the Eight Queens problem."""
    board = QueenBoard()
    solve_queens(board, 0)

if __name__ == "__main__":
    main_eight_queens()

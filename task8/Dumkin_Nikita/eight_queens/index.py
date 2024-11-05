N = 8
board = [[0] * N for _ in range(N)]

def print_board():
    """Prints the board with queens placed on it."""

    for row in board:
        print(" ".join("Q" if col == 1 else "." for col in row))
    print()

def is_safe(row, col):
    """Checks if it's safe to place a queen at (row, col)."""

    for i in range(col):
        if board[row][i] == 1:
            return False

    for i, j in zip(range(row, -1, -1), range(col, -1, -1)):
        if board[i][j] == 1:
            return False

    for i, j in zip(range(row, N), range(col, -1, -1)):
        if i < N and board[i][j] == 1:
            return False

    return True

def solve(col):
    """Recursively tries to place queens in each column."""

    if col >= N:
        return True

    for i in range(N):
        if is_safe(i, col):
            board[i][col] = 1

            if solve(col + 1):
                return True

            board[i][col] = 0

    return False

def main():
    """Main function to solve the 8 Queens problem."""

    if solve(0):
        print("Solution found:")
        print_board()
    else:
        print("No solution exists")

main()
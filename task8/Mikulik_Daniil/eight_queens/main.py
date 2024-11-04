def pipe(*functions):
    def apply_pipe(data):
        for function in functions:
            data = function(data)
        return data
    return apply_pipe


def queens(data):
    yield from place_queen(0, [])


def place_queen(row, board):
    if row == 8:
        yield board
    else:
        for column in range(8):
            if is_safe(row, column, board):
                yield from place_queen(row + 1, board + [column])


def is_safe(row, column, board):
    for r, c in enumerate(board):
        if c == column or \
           c - column == r - row or \
           c - column == row - r:
            return False
    return True


def format_board(solutions):
    for solution in solutions:
        board = []
        for col in solution:
            line = ['.'] * 8
            line[col] = 'Q'
            board.append(''.join(line))
        yield board


def eight_queens_solutions():
    pipeline = pipe(
        queens,
        format_board
    )
    return list(pipeline(None))

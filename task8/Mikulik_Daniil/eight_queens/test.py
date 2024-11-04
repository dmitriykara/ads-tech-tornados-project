from main import eight_queens_solutions
import pytest


def test_eight_queens_solution_count():
    solutions = eight_queens_solutions()
    assert len(solutions) == 92


def test_eight_queens_unique_boards():
    solutions = eight_queens_solutions()
    for board in solutions:
        assert all(row.count('Q') == 1 for row in board)
        cols = [row.index('Q') for row in board]
        assert len(set(cols)) == 8


def test_eight_queens_board_format():
    solutions = eight_queens_solutions()
    for board in solutions:
        assert len(board) == 8
        for row in board:
            assert len(row) == 8


if __name__ == '__main__':
    pytest.main()

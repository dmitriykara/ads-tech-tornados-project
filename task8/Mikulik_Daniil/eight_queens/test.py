from main import is_safe


def test_is_safe():
    board = [-1, -1, -1, -1, -1, -1, -1, -1]
    assert is_safe(board, 0, 0) is True
    board[0] = 0
    assert is_safe(board, 1, 0) is False
    assert is_safe(board, 1, 1) is False
    assert is_safe(board, 1, 2) is True


if __name__ == "__main__":
    unittest.main()

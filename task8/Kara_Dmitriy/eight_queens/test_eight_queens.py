import unittest
from eight_queens import is_safe

class TestEightQueens(unittest.TestCase):
    def test_is_safe(self):
        board = [-1, -1, -1, -1, -1, -1, -1, -1]
        self.assertTrue(is_safe(board, 0, 0))
        board[0] = 0
        self.assertFalse(is_safe(board, 1, 0))
        self.assertFalse(is_safe(board, 1, 1))
        self.assertTrue(is_safe(board, 1, 2))

if __name__ == "__main__":
    unittest.main()

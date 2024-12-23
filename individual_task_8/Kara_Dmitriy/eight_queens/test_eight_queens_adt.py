import unittest
from eight_queens_adt import QueenBoard

class TestEightQueensADT(unittest.TestCase):
    def test_is_safe(self):
        board = QueenBoard()
        self.assertTrue(board.is_safe(0, 0))
        board.place_queen(0, 0)
        self.assertFalse(board.is_safe(1, 0))
        self.assertFalse(board.is_safe(1, 1))
        self.assertTrue(board.is_safe(1, 2))

if __name__ == "__main__":
    unittest.main()

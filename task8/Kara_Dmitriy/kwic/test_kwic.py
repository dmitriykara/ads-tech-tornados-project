import unittest
from kwic import generate_shifts, alphabetize_shifts

class TestKWIC(unittest.TestCase):
    def test_generate_shifts(self):
        lines = ["The quick brown fox"]
        expected_shifts = [
            "The quick brown fox",
            "quick brown fox The",
            "brown fox The quick",
            "fox The quick brown"
        ]
        self.assertEqual(generate_shifts(lines), expected_shifts)

    def test_alphabetize_shifts(self):
        shifts = [
            "quick brown fox The",
            "The quick brown fox",
            "brown fox The quick",
            "fox The quick brown"
        ]
        expected_sorted_shifts = [
            "The quick brown fox",
            "brown fox The quick",
            "fox The quick brown",
            "quick brown fox The"
        ]
        self.assertEqual(alphabetize_shifts(shifts), expected_sorted_shifts)

if __name__ == "__main__":
    unittest.main()

"""Tests for testing basics -- Python equivalent of add_test.go."""

import unittest
from main import add, abs_val, is_even


class TestAdd(unittest.TestCase):
    def test_positive(self):
        self.assertEqual(add(2, 3), 5)

    def test_zeros(self):
        self.assertEqual(add(0, 0), 0)

    def test_negative(self):
        self.assertEqual(add(-1, 1), 0)
        self.assertEqual(add(-3, -7), -10)


class TestAbs(unittest.TestCase):
    def test_negative(self):
        self.assertEqual(abs_val(-7), 7)

    def test_positive(self):
        self.assertEqual(abs_val(5), 5)

    def test_zero(self):
        self.assertEqual(abs_val(0), 0)


class TestIsEven(unittest.TestCase):
    def test_cases(self):
        cases = [
            (0, True),
            (1, False),
            (2, True),
            (-3, False),
            (-4, True),
        ]
        for n, want in cases:
            with self.subTest(n=n):
                self.assertEqual(is_even(n), want)


if __name__ == "__main__":
    unittest.main()

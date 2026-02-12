"""Table-driven tests -- Python equivalent of parse_test.go."""

import unittest
from main import parse_int_safe, clamp


class TestParseIntSafe(unittest.TestCase):
    def test_table(self):
        cases = [
            ("valid positive", "42", 0, 42, False),
            ("valid negative", "-7", 0, -7, False),
            ("valid zero", "0", 99, 0, False),
            ("empty string", "", -1, -1, True),
            ("letters", "abc", -1, -1, True),
            ("float string", "3.14", 0, 0, True),
        ]

        for name, inp, fallback, want, want_err in cases:
            with self.subTest(name=name):
                got, err = parse_int_safe(inp, fallback)
                if want_err:
                    self.assertIsNotNone(err, f"expected error for {inp!r}")
                else:
                    self.assertIsNone(err, f"unexpected error for {inp!r}: {err}")
                self.assertEqual(got, want)


class TestClamp(unittest.TestCase):
    def test_table(self):
        cases = [
            ("in range", 5, 0, 10, 5),
            ("below min", -3, 0, 10, 0),
            ("above max", 15, 0, 10, 10),
            ("at min", 0, 0, 10, 0),
            ("at max", 10, 0, 10, 10),
            ("negative range", -5, -10, -1, -5),
        ]

        for name, n, lo, hi, want in cases:
            with self.subTest(name=name):
                self.assertEqual(clamp(n, lo, hi), want)


if __name__ == "__main__":
    unittest.main()

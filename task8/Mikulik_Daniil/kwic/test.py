from main import kwic_index


def test_kwic_multiple_occurrences():
    text = """Python is versatile.
    Python can be used for scripting.
    Python is also popular in data science."""
    expected = [
        "Python can be used for scripting.",
        "Python is also popular in data science.",
        "Python is versatile.",
        "also popular in data science. Python is",
        "be used for scripting. Python can",
        "can be used for scripting. Python",
        "data science. Python is also popular in",
        "for scripting. Python can be used",
        "in data science. Python is also popular",
        "is also popular in data science. Python",
        "is versatile. Python",
        "popular in data science. Python is also",
        "science. Python is also popular in data",
        "scripting. Python can be used for",
        "used for scripting. Python can be",
        "versatile. Python is"
    ]

    assert kwic_index(text) == expected


def test_kwic_no_occurrence():
    text = "There is no mention of the keyword here."
    expected = [
        "There is no mention of the keyword here.",
        "here. There is no mention of the keyword",
        "is no mention of the keyword here. There",
        "keyword here. There is no mention of the",
        "mention of the keyword here. There is no",
        "no mention of the keyword here. There is",
        "of the keyword here. There is no mention",
        "the keyword here. There is no mention of"
    ]
    assert kwic_index(text) == expected


def test_kwic_case_insensitivity():
    text = "python is great. I love Python."
    expected = [
        "I love Python. python is great.",
        "Python. python is great. I love",
        "great. I love Python. python is",
        "is great. I love Python. python",
        "love Python. python is great. I",
        "python is great. I love Python."
    ]
    assert kwic_index(text) == expected


if __name__ == '__main__':
    import pytest
    pytest.main()

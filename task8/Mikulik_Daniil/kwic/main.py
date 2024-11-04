def pipe(*functions):
    def apply_pipe(data):
        for function in functions:
            data = function(data)
        return data
    return apply_pipe


def read_lines(text):
    for line in text.splitlines():
        yield line


def split_words(lines):
    for line in lines:
        yield line.split()


def generate_shifts(word_lists):
    for words in word_lists:
        for i in range(len(words)):
            # Yield one shift (rotation) at a time
            shift = words[i:] + words[:i]
            yield shift


def alphabetize_shifts(shifts):
    sorted_shifts = sorted(' '.join(shift) for shift in shifts)
    for shift in sorted_shifts:
        yield shift


def kwic_index(text):
    pipeline = pipe(
        read_lines,
        split_words,
        generate_shifts,
        alphabetize_shifts
    )

    return list(pipeline(text))

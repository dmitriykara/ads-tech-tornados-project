from typing import List, Callable

class Event:
    def __init__(self):
        self.subscribers = []

    def subscribe(self, subscriber: Callable):
        self.subscribers.append(subscriber)

    def notify(self, data):
        for subscriber in self.subscribers:
            subscriber(data)


# System Components
class KWICSystem:
    def __init__(self):
        self.input_received_event = Event()
        self.circular_shift_event = Event()
        self.alphabetize_event = Event()

        self.input_received_event.subscribe(self.circular_shift)
        self.circular_shift_event.subscribe(self.alphabetize)
        self.alphabetize_event.subscribe(self.display_results)

        self.shifts = []
        self.results = []

    def input_text(self, lines: List[str]):
        """Receive input and notify subscribers"""
        self.input_received_event.notify(lines)

    def circular_shift(self, lines: List[str]):
        """Generate all circular shifts and notify the next event"""
        self.shifts = []
        for line in lines:
            words = line.split()
            for i in range(len(words)):
                shifted = words[i:] + words[:i]
                self.shifts.append(" ".join(shifted))
        self.circular_shift_event.notify(self.shifts)

    def alphabetize(self, shifts: List[str]):
        """Sort all shifts lexicographically and notify the next event"""
        self.results = sorted(shifts)
        self.alphabetize_event.notify(self.results)

    def display_results(self, results: List[str]):
        """Display the sorted shifts"""
        print("KWIC Index:")
        for result in results:
            print(result)

if __name__ == "__main__":
    kwic = KWICSystem()
    text = [
        "This is the sample text",
        "change it if you want to"
    ]
    kwic.input_text(text)
import random


class ConvNet:
    def __init__(self) -> None:
        pass

    def save(self, path: str) -> None:
        pass

    def load(self, path: str) -> None:
        pass

    def forward(self, X: any) -> any:
        print(X)

        classes = [
            "apple",
            "arm",
            "axe",
            "banana",
            "bed",
            "bee",
            "car",
            "coffee-cup",
            "cookie",
            "donut",
            "door",
            "ear",
            "eye",
            "face",
            "hand",
        ]

        rand_idx = random.randint(0, len(classes) - 1)

        return classes[rand_idx]

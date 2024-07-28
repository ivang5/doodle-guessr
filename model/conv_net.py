from __future__ import annotations
import torch
import torch.nn as nn
from torch.types import Number


class ConvNet(nn.Sequential):
    def __init__(self) -> None:
        super().__init__()
        self.add_module(
            "conv1", nn.Conv2d(in_channels=1, out_channels=32, kernel_size=5, padding=2)
        )
        self.add_module("relu1", nn.ReLU())
        self.add_module("pool1", nn.MaxPool2d(kernel_size=2))
        self.add_module(
            "conv2",
            nn.Conv2d(in_channels=32, out_channels=64, kernel_size=5, padding=2),
        )
        self.add_module("relu2", nn.ReLU())
        self.add_module("pool2", nn.MaxPool2d(kernel_size=2))

        self.add_module("flatten", nn.Flatten())

        self.add_module("fc1", nn.Linear(16384, 2048))
        self.add_module("relu3", nn.ReLU())
        self.add_module("dropout", nn.Dropout(p=0.5))

        self.add_module("fc2", nn.Linear(2048, 5))

    def predict(self, X: torch.Tensor) -> tuple[list[float], Number]:
        X = torch.tensor(X).view((1, 1, 64, 64)).float()
        pred = self(X)
        pred_sm: list[float] = pred[0].softmax(-1).tolist()
        pred_class_id = torch.argmax(pred, dim=1).item()

        return pred_sm, pred_class_id

    def save(self, path: str) -> None:
        torch.save(self, path)

    @staticmethod
    def load(path: str) -> ConvNet:
        return torch.load(path)

import os
import torch
import torch.nn as nn
import json
import random
from torch.utils.data import Subset
from torch.utils.data import Dataset, DataLoader

from conv_net import ConvNet

f = open("./dataset/everything.json")
data = json.load(f)
data_list = [
    (torch.tensor(d["pixels"], dtype=torch.float32).unsqueeze(0), d["id"]) for d in data
]
random.shuffle(data_list)

train_data, valid_data, test_data = (
    data_list[:40000],
    data_list[40000:47000],
    data_list[47000:],
)


class DoodleDataset(Dataset):
    def __init__(self, data_list):
        self.data_list = data_list

    def __len__(self):
        return len(self.data_list)

    def __getitem__(self, idx):
        features, label = self.data_list[idx]
        return features, label


train_dataset = DoodleDataset(train_data)
valid_dataset = DoodleDataset(valid_data)
test_dataset = DoodleDataset(test_data)

batch_size = 64
torch.manual_seed(1)
train_dl = DataLoader(train_dataset, batch_size, shuffle=True)
valid_dl = DataLoader(valid_dataset, batch_size, shuffle=False)
test_dl = DataLoader(test_dataset, batch_size, shuffle=False)

# Load model if exists
model_path = "models/doodle-conv-net.ph"
if os.path.exists(model_path):
    model = ConvNet.load(model_path)
else:
    model = ConvNet()

device = torch.device("cpu")

model = model.to(device)

optimizer = torch.optim.Adam(model.parameters(), lr=0.001)
loss_fn = torch.nn.CrossEntropyLoss()


def train(model: ConvNet, num_epochs, train_dl, valid_dl):
    loss_hist_train = [0] * num_epochs
    accuracy_hist_train = [0] * num_epochs
    loss_hist_valid = [0] * num_epochs
    accuracy_hist_valid = [0] * num_epochs
    for epoch in range(num_epochs):
        model.train()
        for x_batch, y_batch in train_dl:
            x_batch = x_batch.to(device)
            y_batch = y_batch.to(device)
            pred = model(x_batch)
            loss = loss_fn(pred, y_batch)
            loss.backward()
            optimizer.step()
            optimizer.zero_grad()
            loss_hist_train[epoch] += loss.item() * y_batch.size(0)
            is_correct = (torch.argmax(pred, dim=1) == y_batch).float()
            accuracy_hist_train[epoch] += is_correct.sum().cpu()

        loss_hist_train[epoch] /= len(train_dl.dataset)
        accuracy_hist_train[epoch] /= len(train_dl.dataset)

        model.eval()
        with torch.no_grad():
            for x_batch, y_batch in valid_dl:
                x_batch = x_batch.to(device)
                y_batch = y_batch.to(device)
                pred = model(x_batch)
                loss = loss_fn(pred, y_batch)
                loss_hist_valid[epoch] += loss.item() * y_batch.size(0)
                is_correct = (torch.argmax(pred, dim=1) == y_batch).float()
                accuracy_hist_valid[epoch] += is_correct.sum().cpu()

        loss_hist_valid[epoch] /= len(valid_dl.dataset)
        accuracy_hist_valid[epoch] /= len(valid_dl.dataset)

        print(
            f"Epoch {epoch+1} accuracy: {accuracy_hist_train[epoch]:.4f} val_accuracy: {accuracy_hist_valid[epoch]:.4f}"
        )
    return loss_hist_train, loss_hist_valid, accuracy_hist_train, accuracy_hist_valid


torch.manual_seed(1)
num_epochs = 10
hist = train(model, num_epochs, train_dl, valid_dl)

if not os.path.exists("models"):
    os.mkdir("models")

path = "models/doodle-conv-net.ph"
torch.save(model, path)

# path = "models/doodle-conv-net.ph"
# model = torch.load(path)

all_predictions = []
all_labels = []

for batch in test_dl:
    features, labels = batch
    predictions = model(features)
    all_predictions.append(predictions)
    all_labels.append(labels)

all_predictions = torch.cat(all_predictions)
all_labels = torch.cat(all_labels)
is_correct = (torch.argmax(all_predictions, dim=1) == all_labels).float()
accuracy = is_correct.mean().item()
print(f"Accuracy: {accuracy * 100:.2f}%")

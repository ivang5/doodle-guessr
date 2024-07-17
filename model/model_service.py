from flask import Flask, request, jsonify
from test import ConvNet

app = Flask(__name__)

model = ConvNet()
model.load("path/to/model")


@app.post("/infer")
def predict():
    data = request.get_json()
    pixel_array = data.get("pixelArray", [])

    pred = model.forward(pixel_array)

    return jsonify({"prediction": pred})


if __name__ == "__main__":
    app.run(debug=True, host="localhost", port=3001)

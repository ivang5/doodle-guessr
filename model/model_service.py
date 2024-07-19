from flask import Flask, request, jsonify
from conv_net import ConvNet

app = Flask(__name__)

model = ConvNet.load("models/doodle-conv-net.ph")

id_to_class = {
    0: "apple",
    1: "arm",
    2: "axe",
    3: "banana",
    4: "bed",
    5: "bee",
    6: "car",
    7: "coffee cup",
    8: "cookie",
    9: "donut",
    10: "door",
    11: "ear",
    12: "eye",
    13: "face",
    14: "hand",
}


@app.post("/infer")
def predict():
    data = request.get_json()
    pixel_array = data.get("pixelArray", [])

    probs, pred_class_id = model.infer(pixel_array)

    for p, [k, v] in zip(probs, id_to_class.items()):
        perc = p * 100
        print(f"{k}. {v}: {perc:.4f}%")

    return jsonify({"prediction": id_to_class[pred_class_id]})


if __name__ == "__main__":
    app.run(debug=True, host="localhost", port=3001)

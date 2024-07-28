import flask
from conv_net import ConvNet

app = flask.Flask(__name__)

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


@app.post("/predict")
def predict():
    data = flask.request.get_json()
    pixel_array = data.get("pixels", [])

    probs, pred_class_id = model.predict(pixel_array)

    for p, [k, v] in zip(probs, id_to_class.items()):
        perc = p * 100
        print(f"{k}. {v}: {perc:.4f}%")

    return flask.jsonify(
        {
            "prediction": id_to_class[pred_class_id],
            "certainty": probs[pred_class_id],
        }
    )


if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=3001)

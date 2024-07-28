const overlay = document.getElementById("overlay");
const overlayFooter = document.getElementById("overlayFooter");
const chosenWord = document.getElementById("word");
const subText = document.getElementById("subText");
const timeLeft = document.getElementById("timeLeft");
const timer = document.getElementById("timer");
const endGame = document.getElementById("endGame");
const timeScore = document.getElementById("timeScore");
const wordsScore = document.getElementById("wordsScore");
const wordLabel = document.getElementById("wordLabel");
const btnPlayAgain = document.getElementById("startBtn");
const canvas = document.getElementById("canvas");
const wrapper = document.getElementById("container");
const btnClr = document.getElementById("btnClear");
const ctx = canvas.getContext("2d");

let words = [
  "apple",
  "arm",
  "axe",
  "banana",
  "bed",
  // "bee",
  // "car",
  // "coffee cup",
  // "cookie",
  // "donut",
  // "door",
  // "ear",
  // "eye",
  // "face",
  // "hand",
];
let wordsCopy;
let word;
let countdown = 20;
let interval;
let drawing = false;
let lastPos = { x: 0, y: 0 };

const width = this.innerWidth;
const height = this.innerHeight;
const size = Math.min(width, height) * 0.7;

canvas.setAttribute("width", size);
canvas.setAttribute("height", size);

ctx.lineWidth = 12;
ctx.lineCap = "round";

window.onload = () => {
  startNextStage(true);
};

canvas.onmousedown = (e) => {
  drawing = true;
  const currPos = { x: e.offsetX, y: e.offsetY };

  drawDot(currPos.x, currPos.y);
  lastPos = currPos;

  schedulePredictRequest();
};

canvas.onmouseenter = (e) => {
  lastPos = { x: e.offsetX, y: e.offsetY };
};

canvas.onmouseup = () => {
  drawing = false;
};

wrapper.onmouseup = () => {
  drawing = false;
};

canvas.onmousemove = (e) => {
  if (!drawing) return;

  const currPos = { x: e.offsetX, y: e.offsetY };

  drawLine(lastPos.x, lastPos.y, currPos.x, currPos.y);
  lastPos = currPos;
};

btnClr.onclick = () => {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
};

btnPlayAgain.onclick = () => {
  countdown = 20;
  startNextStage(true);
};

const drawLine = (x1, y1, x2, y2) => {
  ctx.beginPath();
  ctx.moveTo(x1, y1);
  ctx.lineTo(x2, y2);
  ctx.stroke();
};

const drawDot = (x, y) => {
  ctx.beginPath();
  ctx.arc(x, y, ctx.lineWidth / 2, 0, Math.PI * 2, true);
  ctx.fill();
};

const getPixelsFromCanvas = (dims) => {
  const offScreenCanvas = document.createElement("canvas");
  const offScreenCtx = offScreenCanvas.getContext("2d");
  const x = dims.x;
  const y = dims.y;

  offScreenCanvas.width = x;
  offScreenCanvas.height = y;

  offScreenCtx.drawImage(canvas, 0, 0, x, y);

  const imageData = offScreenCtx.getImageData(0, 0, x, y);
  const pixels = imageData.data;

  const pixelArray = [];
  for (let y = 0; y < offScreenCanvas.height; y++) {
    for (let x = 0; x < offScreenCanvas.width; x++) {
      const index = (y * offScreenCanvas.width + x) * 4;
      const alpha = pixels[index + 3];

      pixelArray.push(alpha === 0 ? 1 : 0);
    }
  }

  return pixelArray;
};

const requestPrint = async (pixelArray) => {
  const response = await fetch("http://localhost:3000/api/print", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      pixels: pixelArray,
    }),
  });
  let data;
  try {
    data = await response.json();
  } catch (err) {
    data = null;
  }
  return { response, data };
};

const requestPredict = async (pixelArray) => {
  const response = await fetch("http://localhost:3000/api/predict", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      pixels: pixelArray,
    }),
  });
  let data;
  try {
    data = await response.json();
  } catch (err) {
    data = null;
  }
  return data;
};

const startNextStage = (isFirst = false) => {
  if (isFirst) {
    wordsCopy = [...words];
    console.log(wordsCopy);
    chosenWord.classList.remove("word--sm");
  }

  if (wordsCopy.length === 0) {
    return showFinalScore();
  }

  const wordIndex = getRandomIndex(wordsCopy.length - 1);
  word = wordsCopy.splice(wordIndex, 1)[0];
  chosenWord.textContent = word;
  drawing = false;
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  setTimerText();

  if (!isFirst) {
    overlayFooter.classList.add("overlay-footer--hidden");
    timeLeft.textContent = `You have ${countdown} seconds left...`;
    overlay.classList.remove("overlay--hidden");
  } else {
    timeLeft.textContent = "You have 20 seconds...";
    timeLeft.classList.remove("hidden");
    subText.classList.remove("hidden");
    endGame.classList.add("hidden");
  }

  setTimeout(() => {
    overlay.classList.add("overlay--hidden");
  }, 3000);

  wordLabel.textContent = word;
  setTimeout(() => {
    overlayFooter.classList.remove("overlay-footer--hidden");
    interval = setInterval(() => {
      if (countdown <= 0) {
        clearInterval(interval);
        showFinalScore();
      } else {
        countdown--;
        setTimerText();
      }
    }, 1000);
  }, 3500);
};

const showFinalScore = () => {
  overlayFooter.classList.add("overlay-footer--hidden");
  endGame.classList.remove("hidden");
  chosenWord.textContent = "Game over";
  chosenWord.classList.add("word--sm");
  timeScore.textContent = `Time left: ${countdown} seconds`;
  const successfulWords =
    wordsCopy.length === 0 ? words.length : words.length - wordsCopy.length;
  wordsScore.textContent = `Successful words: ${successfulWords}`;
  timeLeft.classList.add("hidden");
  subText.classList.add("hidden");
  overlay.classList.remove("overlay--hidden");
  drawing = false;
};

const getRandomIndex = (max) => {
  return Math.floor(Math.random() * max);
};

const setTimerText = () => {
  timer.textContent =
    countdown.toString().length === 2 ? `00:${countdown}` : `00:0${countdown}`;
};

const schedulePredictRequest = () => {
  timeout = setTimeout(async () => {
    const pixelArray = getPixelsFromCanvas({ x: 64, y: 64 });
    const data = await requestPredict(pixelArray);

    if (data) {
      console.log(data);
      pred = data["prediction"];
      certainty = data["certainty"];

      if (pred && certainty) {
        if (pred == word && certainty > 0.9) {
          clearInterval(interval);
          countdown += 5;
          startNextStage();
        }
      } else {
        console.log("error: pred or certainty is undefined");
      }
    } else {
      console.log("error: data is undefined");
    }

    if (drawing) schedulePredictRequest();
  }, 500);
};

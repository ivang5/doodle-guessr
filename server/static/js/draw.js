const canvas = document.getElementById("canvas");
const wrapper = document.getElementById("container");
const btnClr = document.getElementById("btnClear");
const btnPrint = document.getElementById("btnPrint");
const btnPredict = document.getElementById("btnPredict");
const ctx = canvas.getContext("2d");

let drawing = false;
let lastPos = { x: 0, y: 0 };

const width = this.innerWidth;
const height = this.innerHeight;
const size = Math.min(width, height) * 0.8;

canvas.setAttribute("width", size);
canvas.setAttribute("height", size);

ctx.lineWidth = 12;
ctx.lineCap = "round";

canvas.onmousedown = (e) => {
  drawing = true;
  const currPos = { x: e.offsetX, y: e.offsetY };

  drawDot(currPos.x, currPos.y);
  lastPos = currPos;
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

btnPrint.onclick = () => {
  const pixelArray = getPixelsFromCanvas({ x: 64, y: 64 });

  requestPrint(pixelArray);
};

btnPredict.onclick = () => {
  const pixelArray = getPixelsFromCanvas({ x: 64, y: 64 });

  requestPredict(pixelArray);
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
  return { response, data };
};

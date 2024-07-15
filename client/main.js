const canvas = document.getElementById("canvas");
const wrapper = document.getElementById("container");
const btnClr = document.getElementById("btnClear");
const btnSend = document.getElementById("btnSend");
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

btnSend.onclick = () => {
  const offScreenCanvas = document.createElement("canvas");
  const offScreenContext = offScreenCanvas.getContext("2d");

  offScreenCanvas.width = 28;
  offScreenCanvas.height = 28;

  offScreenContext.drawImage(canvas, 0, 0, 28, 28);

  // const dataB64 = offScreenCanvas.toDataURL("image/png");

  offScreenCanvas.toBlob((blob) => {
    fetch("http://localhost:6969", {
      method: "POST",
      body: blob,
    })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  });
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

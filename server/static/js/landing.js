const startBtn = document.getElementById("startBtn");
const testSetBtn = document.getElementById("testSetBtn");
const testReadBtn = document.getElementById("testReadBtn");

startBtn.onclick = () => {
    window.location.href = "/draw.html";
};

testSetBtn.onclick = async () => {
    const response = await fetch("http://localhost:3000/api/scores", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            name: "prki",
            points: 35,
        }),
    });

    console.log(response);
};

testReadBtn.onclick = async () => {
    const response = await fetch("http://localhost:3000/api/scores", {
        method: "GET",
    });

    console.log(response);
};

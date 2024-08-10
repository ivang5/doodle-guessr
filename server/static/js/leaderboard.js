const leaderboard = document.getElementById("leaderboard");

window.onload = async () => {
  const scores = await getLeaderboard();

  if (scores) {
    scores.scores.forEach((score, index) => {
      const leaderboardItem = document.createElement("li");
      const itemName = document.createElement("span");
      const itemScore = document.createElement("span");

      leaderboardItem.className = "list-item";
      itemName.appendChild(
        document.createTextNode(`${index + 1}. ${score.name}`)
      );
      itemScore.appendChild(document.createTextNode(score.points));
      leaderboardItem.appendChild(itemName);
      leaderboardItem.appendChild(itemScore);
      leaderboard.appendChild(leaderboardItem);
    });
  }
};

const getLeaderboard = async () => {
  const response = await fetch("http://localhost:3000/api/scores", {
    method: "GET",
  });

  let data = null;
  try {
    data = await response.json();
  } catch (err) {
    console.log(err);
  }

  return data;
};

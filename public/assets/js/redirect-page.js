const username = window.location.href.split('/').pop();
const statsUrl = `https://lc.andrewpavlov.org/api/slug/${username}/stats`;
const leetCodeProfileUrl = `https://lc.andrewpavlov.org/${username}/redirect`;

function redirectToMain() {
  window.location.href = "https://lc.andrewpavlov.org/";
}

document
  .getElementById("redirectToLeetCode").addEventListener("click", function () {
    window.location.href = leetCodeProfileUrl;
  });

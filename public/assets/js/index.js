// Form submission handler
document.getElementById('usernameForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    const username = document.getElementById('usernameInput').value.trim().toLowerCase();
    if (username === "") {
        alert("Please enter your LeetCode username.");
        return;
    }

    // Generate the badge URL and redirect URL
    const badgeUrl = `https://lc.andrewpavlov.org/api/slug/${username}/badge.svg`;
    const redirectUrl = `https://lc.andrewpavlov.org/redirect-page/${username}`;

    // Generate Markdown embed code
    document.getElementById('markdownEmbed').textContent = `[![LeetCode Badge](${badgeUrl})](${redirectUrl})`;

    // Generate HTML embed code
    document.getElementById('htmlEmbed').textContent = `<a href="${redirectUrl}" target="_blank"><img src="${badgeUrl}" alt="LeetCode Badge"></a>`;

    // Display the container with the embed codes
    document.getElementById('badgeLinks').style.display = 'block';
});

// Function to copy text to clipboard
function copyToClipboard(elementId) {
    const text = document.getElementById(elementId).textContent;

    // Create a temporary textarea element to hold the text
    const tempInput = document.createElement('textarea');
    tempInput.value = text; // Set the value of the textarea to the text
    document.body.appendChild(tempInput); // Append the textarea to the DOM
    tempInput.select(); // Select the text inside the textarea
    document.execCommand('copy'); // Copy the selected text to clipboard
    document.body.removeChild(tempInput); // Remove the textarea from the DOM

    showCopyNotification(); // Show the notification for successful copy
}

// Function to show the copy notification
function showCopyNotification() {
const notification = document.getElementById('copyNotification');
notification.style.display = 'block'; 
notification.style.opacity = '1'; 
notification.style.transform = 'translateY(0)';

setTimeout(() => {
    notification.style.opacity = '0'; 
    notification.style.transform = 'translateY(-20px)'; 
    setTimeout(() => {
        notification.style.display = 'none';
    }, 300); 
}, 2000); 
}
async function send_msg() {
  const message = document.getElementById('message-input').value;
  console.log(message);
  if (message) {
    document.getElementById('message-input').value = "";
    var username = document.getElementById("uname");
    username = username.textContent;
    var time = new Date().toLocaleTimeString();
    document.getElementById('messages').innerHTML += `
      <div class="notification-item" style="color:var(--notification-item-user-text-color);">
        ${username}:${message}
        <span class="timestamp">${time}</span>
        </div>`;

    var m = document.getElementById("messages");
    m.scrollTop = m.scrollHeight;
    const response = await fetch('/ai_chat', {
      method: 'POST',
      body: JSON.stringify({
        message: message,
        username: username
      })
    });
    const data = await response.json();
    console.log(data);
    document.getElementById('messages').innerHTML += `
      <div class="notification-item" style="color:var(--notification-item-system-text-color);">
        ${data.username}:${data.massage}
        <span class="timestamp">${data.time}</span>
        </div>`;

    var m = document.getElementById("messages");
    m.scrollTop = m.scrollHeight;
  }
}
var message_form = document.getElementById("message_form")
message_form.addEventListener("submit", function (event) {
  event.preventDefault();
  send_msg();
})
var msg_list = document.getElementById("messages");
msg_list.scrollTop = msg_list.scrollHeight;
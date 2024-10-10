
get_msg = function () {
  fetch('/query_msg', {
    method: 'GET',
  })
    .then(response => response.json())
    .then(data => {
      if (data.parmcode == 1) {
        for (var i = 0; i < data.msgs.length; i++) {

          // data.msgs[i].time = data.msgs[i].time.split(".")[0];
          data.msgs[i].time = String(data.msgs[i].time).match(/[^:]*:[^:]*/);
          document.getElementById('messages').innerHTML += `
                              <div class="notification-item" style="color:var(--notification-item-user-text-color);">
                                  ${data.msgs[i].username}:${data.msgs[i].message}
                                  <span class="timestamp">${data.msgs[i].time}</span>
                                  </div>`;

        }
      } else {
        alert("获取失败，请重试")
      }
      var m = document.getElementById("messages");
      m.scrollTop = m.scrollHeight;
    });
}
function sendMessage() {
  const message = document.getElementById('message-input').value;
  if (message) {
    uname = document.getElementById("uname").innerText;
    console.log(uname)
    fetch('/send_msg', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: uname,
        message: message,
        time: new Date().toLocaleString() + "." + new Date().getMilliseconds()
      })
    })
      .then(response => response.json())
      .then(data => {
        if (data.parmcode == 1) {
          document.getElementById('message-input').value = '';
        } else {
          alert("发送失败，请重试")
        }
      });
  }
}
var message_form = document.getElementById("message_form")
message_form.addEventListener("submit", function (event) {
  event.preventDefault();
  sendMessage();
})


var host = window.location.host
const socket = new WebSocket("ws://" + host + "/ws_receive_msg");
socket.onmessage = function (event) {
  const message = event.data;
  var data = JSON.parse(message);

  data.time = String(data.time).match(/[^:]*:[^:]*/);
  document.getElementById('messages').innerHTML += `
      <div class="notification-item" style="color:var(--notification-item-user-text-color);">
        ${data.username}:${data.message}
        <span class="timestamp">${data.time}</span>
        </div>`;
  var m = document.getElementById("messages");
  m.scrollTop = m.scrollHeight;
};

socket.onerror = function (error) {
  const net_status = document.getElementById("net_status");
  net_status.innerHTML = "网络连接出错";
  net_status.style.backgroundColor = "crimson";

};
socket.onopen = function (event) {
  const net_status = document.getElementById("net_status");
  net_status.innerHTML = "网络已连接";
  net_status.style.backgroundColor = "#28a745";
};
socket.onclose = function (event) {
  const net_status = document.getElementById("net_status");
  net_status.innerHTML = "网络未连接";
  net_status.style.backgroundColor = "crimson";
};
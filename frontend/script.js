const chatDiv = document.getElementById("chat");
var nickname = "";
var socket = null;

function createDivWithClass(className) {
  let element = document.createElement("div");
  element.classList.add(className);
  return element;
}

function messageToDiv(msgJson) {
  let message = createDivWithClass("message");
  let from = createDivWithClass(
    msgJson.author.toLowerCase() == "you" ? "from-user" : "not-from-user"
  );
  let author = createDivWithClass("author");
  let body = createDivWithClass("body");
  let timeSent = createDivWithClass("time-sent");

  author.innerText = msgJson.author;
  body.innerText = msgJson.body;
  timeSent.innerText = msgJson.createdAt;

  message.appendChild(from);
  from.appendChild(author);
  from.appendChild(body);
  from.appendChild(timeSent);

  return message;
}
function getCurrentTime() {
  date = new Date()
  hour = date.getHours();
  minutes = date.getMinutes()
  pm = hour - 12 < 0 ? "PM" : "AM";

  return `${hour}:${minutes} ${pm}`
}
function sendMessageButton() {
  let messageText = document.getElementById("message-input").value;
  let message = {
    author: "you",
    body: messageText,
    createdAt: getCurrentTime(),
  };

  sendMessageToServer(message);
  addMessageToChat(message);
}

function sendMessageToServer(message) {
  socket.send(JSON.stringify(message));
}

function addMessageToChat(message) {
  chatDiv.appendChild(messageToDiv(message));
  chatDiv.scrollTop = chatDiv.scrollHeight;
}

function setIsConnected(b) {
  statusElem = document.getElementById("status");
  if (b) {
    statusElem.innerText = "Status: connected";
    return;
  }
  statusElem.innerText = "Status: disconnected";
}

function connectToWebsocket() {
  nickname = document.getElementById("nickname").value;
  if (!nickname) {
    alert("Please enter a nickname");
    return;
  }
  if (!socket) {
    socket = new WebSocket("ws://localhost:8080/chat?nickname=" + nickname);
    socket.onopen = (event) => {
      document.getElementById("");
    };

    socket.onmessage = (event) => {
      msg = JSON.parse(event.data);
      addMessageToChat(msg);
    };

    setIsConnected(true);
  } else {
    socket.close();
    setIsConnected(false);
  }
}

const wsScheme = location.protocol === "https:" ? "wss" : "ws";
const wsUrl = wsScheme + "://" + location.host + location.pathname + "/ws";
const ws = new WebSocket(wsUrl);

const form = document.getElementById("chat");

ws.onclose = () => {
	const messageInputElement = document.getElementById("message");
	messageInputElement.disabled = true;
	messageInputElement.value = "";
	messageInputElement.placeholder = "Chatroom is closed.";
};

ws.onopen = () => {
	const messageInputElement = document.getElementById("message");
	messageInputElement.disabled = false;
	messageInputElement.value = "";
	messageInputElement.placeholder = "Type a message...";
};

ws.onmessage = (e) => {
	const msg = e.data;
	const box = document.getElementById("chat-box");
	const div = document.createElement("div");
	const sender = document.createElement("u");

	sender.textContent = msg.username + ":";
	div.appendChild(sender);

	div.textContent = " " + msg.text;
	box.appendChild(div);

	box.scrollTop = box.scrollHeight;
};

form.onsubmit = async (e) => {
	e.preventDefault();

	// retrieving the message, then sending it to the server
	const input = document.getElementById("message");
	if (input.value.trim() !== "") {
		const message = input.value;
		input.value = "";

		const messageResponse = await fetch("http://localhost:8080/api/lobby/" + localStorage.getItem("lobby_id") + "/message", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				sender_id: localStorage.getItem("client_id"),
				body: message,
			}),
		}).then((response) => response.json());

		// if the message was successfully received and created server-side, we send it to other clients through the websocket
		if (messageResponse.code == 200) {
			ws.send(messageResponse);
			console.log("message sent");
		} else {
			alert("message couldn't be sent!");
		}
	}
};

document.addEventListener("DOMContentLoaded", () => {
	const lobbyId = localStorage.getItem("lobby_id");

	if (!lobbyId) {
		window.location.href = "/";
		return;
	}

	document.getElementById("lobby-id").innerText = lobbyId;
});
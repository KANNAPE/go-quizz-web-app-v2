// const wsScheme = location.protocol === "https:" ? "wss" : "ws";
// const wsUrl = wsScheme + "://" + location.host + location.pathname;
// const ws = new WebSocket(wsUrl);

// ws.onmessage = (e) => {
//     const msg = JSON.parse(e.data);
//     const box = document.getElementById("chat-box");
//     const div = document.createElement("div");
//     const sender = document.createElement("u");

//     sender.textContent = msg.username + ":";
//     div.appendChild(sender);

//     div.textContent = " " + msg.text;
//     box.appendChild(div);

//     box.scrollTop = box.scrollHeight;
// };

// document.getElementById("chat-form").onsubmit = (e) => {
//     e.preventDefault();

//     const input = document.getElementById("chat-input");
//     if (input.value.trim() !== "") {
//         ws.send(input.value);
//         input.value = "";
//     }
// };

document.addEventListener("DOMContentLoaded", () => {
	const lobbyId = localStorage.getItem("lobby_id");

	if (!lobbyId) {
		window.location.href = "/";
		return;
	}

	document.getElementById("lobby-id").innerText = lobbyId;
});
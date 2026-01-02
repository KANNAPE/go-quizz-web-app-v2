var createLobbyForm = document.getElementById("create-lobby-form");
var joinLobbyForm = document.getElementById("join-lobby-form");

async function createLobby(event) {
    event.preventDefault();

    // making a request to the server to create the lobby
    const lobbyCreationResponse = await fetch("http://localhost:8080/api/lobby", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
    }).then((response) => response.json());

    if (lobbyCreationResponse.code !== 200) {
        alert(lobbyCreationResponse.message);
        return;
    }

    const joinLobbyResponse = await fetch("http://localhost:8080/api/lobby/" + lobbyCreationResponse.data.created_lobby_id + "/connect", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            username: createLobbyForm.username.value,
        }),
    }).then((response) => response.json());

    if (joinLobbyResponse.code !== 200) {
        alert("error joining the lobby!");
        return;
    }

    localStorage.setItem("lobby_id", joinLobbyResponse.data.lobby_id);
    localStorage.setItem("client_id", joinLobbyResponse.data.client_id);

    createLobbyForm.submit(); // to get to lobby with the POST method
}

async function joinLobby(event) {
    event.preventDefault();

    const urlParams = new URLSearchParams(window.location.search);
    const lobbyId = urlParams.get("id");

    if (lobbyId === null) {
        window.location.href = "/";
        return;
    }

    const joinLobbyResponse = await fetch("http://localhost:8080/api/lobby/" + lobbyId + "/connect", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            username: joinLobbyForm.username.value,
        }),
    }).then((response) => response.json());

    if (joinLobbyResponse.code !== 200) {
        alert("error joining the lobby!");
        return;
    }

    localStorage.setItem("lobby_id", joinLobbyResponse.data.lobby_id);
    localStorage.setItem("client_id", joinLobbyResponse.data.client_id);

    joinLobbyForm.submit(); // to get to lobby with the POST method
}

joinLobbyForm?.addEventListener("submit", joinLobby);
createLobbyForm?.addEventListener("submit", createLobby);

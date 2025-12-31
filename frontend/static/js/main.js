var Form = document.getElementById("create-lobby-form");

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
            username: Form.username.value,
        }),
    }).then((response) => response.json());

    if (joinLobbyResponse.code !== 200) {
        alert("error joining the lobby!");
        return;
    }

    localStorage.setItem("lobby_id", lobbyCreationResponse.data.created_lobby_id);

    Form.submit(); // to get to lobby with the POST method
}

Form.addEventListener("submit", createLobby);

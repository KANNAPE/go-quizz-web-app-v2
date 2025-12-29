const form = document.getElementById("create-lobby-form");

function createLobby(event) {
    event.preventDefault();

    // making a request to create the lobby
    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function () {
        if (this.readyState == XMLHttpRequest.DONE && this.status == 200) {
            var response = JSON.parse(this.responseText);

            var lobbyId = response.data.created_lobby_id
            console.log(lobbyId);

            joinLobby(lobbyId);
        }
    };

    xhttp.open("POST", "http://localhost:8080/api/lobby", true);
    xhttp.send();
}

function joinLobby(lobbyId) {
    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function () {
        if (this.readyState == XMLHttpRequest.DONE && this.status == 200) {
            var response = JSON.parse(this.responseText);

            console.log(response);

            form.submit();
        } else {
            console.log(this.responseText);
        }
    };

    const formData = new FormData(form);
    const username = formData.get("username")

    xhttp.open("POST", "http://localhost:8080/api/lobby/" + lobbyId + "/connect", true);
    xhttp.setRequestHeader("Content-Type", "application/json");

    xhttp.send(JSON.stringify({
        username: username,
    }));
}

form.addEventListener("submit", createLobby);
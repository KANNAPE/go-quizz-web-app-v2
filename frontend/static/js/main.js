var Form = document.getElementById("create-lobby-form");

function createLobby(event) {
    event.preventDefault();

    // making a request to create the lobby
    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function () {
        if (this.readyState == XMLHttpRequest.DONE) {
            var response = JSON.parse(this.responseText);

            if (this.status == 200) {
                var lobbyId = response.data.created_lobby_id
                console.log(lobbyId);

                joinLobby(lobbyId);
            } else {
                console.log(response.message);
            }
        };

        xhttp.open("POST", "http://localhost:8080/api/lobby", true);
        xhttp.send();
    }

    function joinLobby(lobbyId) {
        var xhttp = new XMLHttpRequest();

        const formData = new FormData(Form);
        const username = formData.get("username");

        xhttp.onreadystatechange = function () {
            if (this.readyState == XMLHttpRequest.DONE) {
                var response = JSON.parse(this.responseText);

                if (this.status == 200) {
                    Form.submit();
                } else {
                    console.log(response.message);
                }
            };

            xhttp.open("POST", "http://localhost:8080/api/lobby/" + lobbyId + "/connect", true);
            xhttp.setRequestHeader("Content-Type", "application/json");

            xhttp.send(JSON.stringify({
                username: username,
            }));
        }
    }
}

Form.addEventListener("submit", createLobby);

let ws;

document.addEventListener("DOMContentLoaded", Connect);

function Connect() {
    ws = new WebSocket("ws://" + window.location.host + "/ws");
    console.log("Websocket - status: " + ws.readyState);

    ws.onopen = function () {
        console.log("Connection opened..." + this.readyState);
    };

    ws.onmessage = function (msg) {
        UpdateLog(msg.data);
    };

    ws.onerror = function (msg) {
        console.log("Error occured sending..." + msg.data);
    };

    ws.onclose = function (msg) {
        alert("Disconnected - status " + this.readyState);
    };
}

function OpenDoor() {
    ws.send(JSON.stringify({
        event: "OpenDoor"
    }));
}

function UpdateLog(jsonMessage) {
    let logBlock = document.getElementById("tableLog");

    let logs = JSON.parse(jsonMessage);

    for (let i = 0; i < logs.length; i++) {

        let logRow = document.getElementById("log" + logs[i].id);

        if (!logRow) {
            let row = document.createElement("tr");
            row.id = "log" + logs[i].id;

            let time = document.createElement("td");
            time.className = "timeLog";
            time.innerHTML = logs[i].time.split("T")[0] + " - " + logs[i].time.substring(11, 19) + " "; // лень думать
            row.appendChild(time);

            let name = document.createElement("td");
            name.innerHTML = logs[i].name;
            row.appendChild(name);

            let event = document.createElement("td");
            event.innerHTML = logs[i].event;
            row.appendChild(event);

            let uuid = document.createElement("td");
            uuid.innerHTML = logs[i].uuid;
            row.appendChild(uuid);

            logBlock.appendChild(row);

            document.getElementById("log").scrollTop = 9999;
        }
    }
}
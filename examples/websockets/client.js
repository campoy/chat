var sock = new WebSocket("ws://localhost:8080/");
sock.onmessage = function(m) { console.log("Received:", m.data); }
sock.send("Hello!\n")
var socket;

const connectToWebSocket = (onMessageCallback) => {
  if (!socket || socket.readyState === WebSocket.CLOSED) {
    socket = new WebSocket("ws://localhost:8000/ws");

    socket.onopen = () => {
      console.log("Successfully Connected");
    };

    socket.onmessage = (msg) => {
      console.log("Message received ", msg);
      onMessageCallback(msg.data)
    };

    socket.onerror = (err) => {
      console.log("error came", err);
    };

    socket.onclose = (event) => {
      console.log("socket closed connection", event);
    };
  } else {
        console.log("WebSocket is already connected or connecting.");
    }
};

const sendMsg = (msg) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connectToWebSocket, sendMsg };

"use client";
import { connectToWebSocket, sendMsg } from "./api";
import { useEffect, useState } from "react";

export default function Home() {
  const [inputMessage, setInputMessage] = useState();
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    connectToWebSocket((newMsg) => {
      setMessages((prevMessages) => [...prevMessages, newMsg]);
    });
  }, []);

  const handleSend = () => {
    console.log("mess", inputMessage);
    sendMsg(inputMessage);
    setInputMessage("");
  };

  return (
    <>
      <div className="App mx-auto flex justify-center mt-12 flex-col items-center">
        <h2 className="text-xl font-bold">Real Time Chat Application</h2>
        <div className="mt-4 flex gap-4">
          <input
            type="text"
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
            placeholder="Type your message here..."
            className="border p-2 rounded-md text-lg w-72"
          />
          <button
            onClick={handleSend}
            className="bg-blue-600 p-2 px-6 text-white font-bold rounded-md text-lg cursor-pointer"
          >
            Send
          </button>
        </div>
      </div>
      <div className="mt-6 mx-auto w-1/2">
        <ul className="list-disc list-inside">
          <h2 className="text-xl font-bold">Messages:</h2>
          {messages && messages.length > 0 ? (
            messages.map((message, index) => (
              <li key={index} className="text-lg">
                {message}
              </li>
            ))
          ) : (
            <p>No messages yet...</p>
          )}
        </ul>
      </div>
    </>
  );
}

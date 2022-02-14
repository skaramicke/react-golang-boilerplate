import React, { useEffect } from "react";
import { render } from "react-dom";

const App = () => {

  const [backendMessage, setBackendMessage] = React.useState("Loading...");

  useEffect(() => {
    fetch("/api/hello")
      .then(res => res.json())
      .then(data => {
        console.log(data.data);
        setBackendMessage(data.data)
      });
  }, []);

  return (
    <div>From Backend: {backendMessage}</div>
  );
}

render(<App />, document.getElementById("app"));
import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

// React.StrictMode만 감싸면 됩니다. Router는 App 컴포넌트 내에서 처리.
const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <App />{" "}
    {/* App 컴포넌트는 이미 Router로 감싸져 있으므로 추가적인 Router는 필요 없습니다. */}
  </React.StrictMode>
);

reportWebVitals();

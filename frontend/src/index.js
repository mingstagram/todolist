import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

// React.StrictMode를 제거하고 App만 렌더링
const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <App /> // React.StrictMode 제거
);

reportWebVitals();

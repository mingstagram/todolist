import React, { useState } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  useNavigate,
  Navigate,
} from "react-router-dom";
import LoginForm from "./components/users/LoginForm";
import SignupForm from "./components/users/SignupForm";
import TodoApp from "./components/todos/TodoApp";
import { saveUsers } from "./api/UsersApi";

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  return (
    <Router>
      {" "}
      {/* App 내에서 Router를 감싸는 부분은 그대로 둡니다 */}
      <Routes>
        <Route
          path="/login"
          element={<LoginForm setIsLoggedIn={setIsLoggedIn} />}
        />
        <Route path="/signup" element={<SignupForm />} />
        <Route
          path="/todos"
          element={isLoggedIn ? <TodoApp /> : <Navigate to="/login" />}
        />
        <Route path="*" element={<Navigate to="/login" />} />
      </Routes>
    </Router>
  );
};

export default App;

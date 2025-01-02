import React, { useState, useEffect } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import LoginForm from "./components/users/LoginForm";
import SignupForm from "./components/users/SignupForm";
import TodoApp from "./components/todos/TodoApp";
import KakaoRedirection from "./api/KakaoRedirection";

const App = () => {
  // 로그인 상태 초기화: localStorage에서 로그인 정보 확인
  const [isLoggedIn, setIsLoggedIn] = useState(() => {
    // localStorage에서 "accessToken"이 있으면 로그인 상태로 설정
    return localStorage.getItem("accessToken") !== null;
  });

  // 로그인 상태가 변경되면 localStorage에 저장
  useEffect(() => {
    if (isLoggedIn) {
      // 로그인 상태일 때 localStorage에 로그인 정보를 저장
      localStorage.setItem("accessToken", "your-token"); // 실제로는 서버에서 받은 토큰을 저장해야 함
    } else {
      // 로그아웃 상태일 때 localStorage에서 로그인 정보 삭제
      localStorage.removeItem("accessToken");
    }
  }, [isLoggedIn]); // isLoggedIn 상태가 변경될 때마다 실행

  return (
    <Router>
      <Routes>
        <Route path="/kakao/callback" element={<KakaoRedirection />} />
        <Route
          path="/login"
          element={<LoginForm setIsLoggedIn={setIsLoggedIn} />}
        />
        <Route path="/signup" element={<SignupForm />} />
        <Route
          path="/todos"
          element={
            isLoggedIn ? (
              <TodoApp setIsLoggedIn={setIsLoggedIn} />
            ) : (
              <Navigate to="/login" />
            )
          }
        />
        <Route path="*" element={<Navigate to="/login" />} />
      </Routes>
    </Router>
  );
};

export default App;

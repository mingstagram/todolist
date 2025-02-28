import React, { useEffect, useState } from "react";
import { createGlobalStyle } from "styled-components";
import TodoTemplate from "./TodoTemplate";
import TodoHead from "./TodoHead";
import TodoList from "./TodoList";
import TodoCreate from "./TodoCreate";
import { countTasksByDate, getTasksByDate } from "../../api/TodosApi";

const GlobalStyle = createGlobalStyle`
  body {
    background: #e9ecef;
  }
`;

function TodoApp({ setIsLoggedIn }) {
  const sessionStorage = window.sessionStorage;
  const userId = sessionStorage.getItem("userId");
  const [tasks, setTasks] = useState([]);
  const [error, setError] = useState(null);
  const [date, setDate] = useState(new Date());
  const [count, setCount] = useState(0);

  const fetchTasksByDate = async (date) => {
    // 날짜 포맷을 명시적으로 설정 (예: YYYY-MM-DD)
    const formattedDate = date.toISOString().split("T")[0]; // ISO 포맷에서 날짜만 추출
    try {
      const data = await getTasksByDate(formattedDate, userId); // 서버 호출
      setTasks(data);
    } catch (err) {
      setError("Failed to load tasks.");
      console.error(err);
    }
  };

  const fetchCountTasksByDate = async (date) => {
    const formattedDate = date.toISOString().split("T")[0]; // ISO 포맷에서 날짜만 추출
    try {
      const data = await countTasksByDate(formattedDate, userId); // 서버 호출
      setCount(data.data.count);
    } catch (err) {
      setError("Failed to load tasks.");
      console.error(err);
    }
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    localStorage.removeItem("accessToken");
    localStorage.removeItem("userId");
  };

  useEffect(() => {
    console.log(date);
    fetchTasksByDate(date);
    fetchCountTasksByDate(date);
  }, [date]);

  return (
    <>
      <GlobalStyle />
      <TodoTemplate>
        <TodoHead
          date={date}
          setDate={setDate}
          count={count}
          refreshTasks={() => fetchTasksByDate(date)}
          handleLogout={handleLogout}
        />
        <TodoList
          tasks={tasks}
          setTasks={setTasks}
          fetchCountTasksByDate={fetchCountTasksByDate}
          date={date}
        />
        <TodoCreate
          fetchTasksByDate={fetchTasksByDate}
          fetchCountTasksByDate={fetchCountTasksByDate}
          date={date}
        />
      </TodoTemplate>
    </>
  );
}

export default TodoApp;

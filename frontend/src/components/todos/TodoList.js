import React from "react";
import styled from "styled-components";
import TodoItem from "./TodoItem";
import { deleteTasks, updateCheckedStatus } from "../../api/TodosApi"; // PUT 요청 API 호출 함수

const TodoListBlock = styled.div`
  flex: 1;
  padding: 20px 32px;
  padding-bottom: 48px;
  overflow-x: auto;
`;

function TodoList({ tasks, setTasks, fetchCountTasksByDate, date }) {
  const handleToggle = async (id, isChecked) => {
    try {
      // 서버로 PUT 요청
      await updateCheckedStatus(id, isChecked);
      // 클라이언트 상태 업데이트 (tasks)
      const updatedTasks = tasks.map((task) =>
        task.id === id ? { ...task, is_checked: isChecked } : task
      );
      setTasks(updatedTasks); // tasks 업데이트
      // 서버에서 count 다시 가져오기
      await fetchCountTasksByDate(date);
    } catch (error) {
      console.error("Failed to update task status:", error);
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm("삭제 하시겠습니까?")) {
      try {
        await deleteTasks(id);
        const deletedTasks = tasks.filter((task) => task.id !== id); // tasks에서 id가 일치하는 task 제거
        setTasks(deletedTasks);
        await fetchCountTasksByDate(date);
      } catch (error) {
        console.error("Failed to update task status:", error);
      }
    }
  };

  return (
    <TodoListBlock>
      {tasks?.map((item) => (
        <TodoItem
          key={item.id}
          id={item.id}
          text={item.task}
          done={item.is_checked}
          onToggle={handleToggle}
          onDelete={handleDelete}
        />
      ))}
    </TodoListBlock>
  );
}

export default TodoList;

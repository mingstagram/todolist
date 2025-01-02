import React, { useEffect, useState } from "react";
import styled, { css } from "styled-components";
import { MdAdd } from "react-icons/md";
import { saveTasks } from "../../api/TodosApi";

const CircleButton = styled.button`
  background: #38d9a9;
  &:hover {
    background: #63e6be;
  }
  &:active {
    background: #20c997;
  }

  z-index: 5;
  cursor: pointer;
  width: 80px;
  height: 80px;
  display: block;
  align-items: center;
  justify-content: center;
  font-size: 60px;
  position: absolute;
  left: 50%;
  bottom: 0px;
  transform: translate(-50%, 50%);
  color: white;
  border-radius: 50%;
  border: none;
  outline: none;
  display: flex;
  align-items: center;
  justify-content: center;

  transition: 0.125s all ease-in;
  ${(props) =>
    props.open &&
    css`
      background: #ff6b6b;
      &:hover {
        background: #ff8787;
      }
      &:active {
        background: #fa5252;
      }
      transform: translate(-50%, 50%) rotate(45deg);
    `}
`;

const InsertFormPositioner = styled.div`
  width: 100%;
  bottom: 0;
  left: 0;
  position: absolute;
`;

const InsertForm = styled.form`
  background: #f8f9fa;
  padding-left: 32px;
  padding-top: 32px;
  padding-right: 32px;
  padding-bottom: 72px;

  border-bottom-left-radius: 16px;
  border-bottom-right-radius: 16px;
  border-top: 1px solid #e9ecef;
`;

const Input = styled.input`
  padding: 12px;
  border-radius: 4px;
  border: 1px solid #dee2e6;
  width: 100%;
  outline: none;
  font-size: 18px;
  box-sizing: border-box;
`;

function TodoCreate({ fetchTasksByDate, date, fetchCountTasksByDate }) {
  const sessionStorage = window.sessionStorage;
  const userId = sessionStorage.getItem("userId");
  const [open, setOpen] = useState(false);
  const [tasks, setTasks] = useState({
    task: "",
    user_id: parseInt(userId, 10),
  });
  const [error, setError] = useState(null);

  const onToggle = () => setOpen(!open);

  const handleInputChange = (e) => {
    setTasks({
      ...tasks,
      task: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!tasks.task.trim()) return; // 공백 입력 방지

    const updatedTask = {
      ...tasks,
      created_at: date,
    };

    try {
      await saveTasks(updatedTask);
      setTasks({
        ...tasks,
        task: "",
      });
      setOpen(false);
      fetchTasksByDate(date);
      await fetchCountTasksByDate(date);
    } catch (error) {
      setError("Failed to save tasks");
      console.error("Error save tasks:", error);
    }
  };

  return (
    <>
      {open && (
        <InsertFormPositioner>
          <InsertForm onSubmit={handleSubmit}>
            <Input
              value={tasks.task}
              onChange={handleInputChange}
              autoFocus
              placeholder="할일을 입력하세요."
            />
          </InsertForm>
        </InsertFormPositioner>
      )}
      <CircleButton onClick={onToggle} open={open}>
        <MdAdd />
      </CircleButton>
    </>
  );
}

export default TodoCreate;

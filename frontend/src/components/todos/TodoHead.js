import React, { useState } from "react";
import styled from "styled-components";

const TodoHeadBlock = styled.div`
  padding: 48px 32px 24px 32px;
  border-bottom: 1px solid #e9ecef;
  display: flex;
  flex-direction: column;
  align-items: center;
  h1 {
    margin: 0px;
    font-size: 36px;
    color: #343a40;
  }
  .day {
    margin-top: 4px;
    color: #868e96;
    font-size: 21px;
  }
  .tasks-left {
    color: #20c997;
    font-size: 18px;
    margin-top: 40px;
    font-weight: bold;
  }
`;

const Navigation = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 20px;

  button {
    background: none;
    border: none;
    color: #343a40;
    font-size: 24px;
    cursor: pointer;
    padding: 0 15px;

    &:hover {
      color: #20c997;
    }

    &:disabled {
      color: #d3d3d3;
      cursor: not-allowed;
    }
  }
`;

const nowDate = (currentDate) => {
  const year = currentDate.getFullYear();
  const month = String(currentDate.getMonth() + 1).padStart(2, "0");
  const day = String(currentDate.getDate()).padStart(2, "0");
  return `${year}년 ${month}월 ${day}일`;
};

const nowDayOfTheWeek = (currentDate) => {
  const daysOfWeek = [
    "일요일",
    "월요일",
    "화요일",
    "수요일",
    "목요일",
    "금요일",
    "토요일",
  ];
  return daysOfWeek[currentDate.getDay()];
};

function TodoHead({ date, setDate, count }) {
  const handlePreviousDay = () => {
    // 새로운 Date 객체를 만들어야 상태가 올바르게 동작
    setDate(new Date(date.getTime() - 24 * 60 * 60 * 1000));
  };

  const handleNextDay = () => {
    // 새로운 Date 객체를 만들어야 상태가 올바르게 동작
    setDate(new Date(date.getTime() + 24 * 60 * 60 * 1000));
  };

  return (
    <TodoHeadBlock>
      <Navigation>
        <button onClick={handlePreviousDay}>&lt;</button>
        <h1>{nowDate(date)}</h1>
        <button onClick={handleNextDay}>&gt;</button>
      </Navigation>
      <div className="day">{nowDayOfTheWeek(date)}</div>
      <div className="tasks-left">할 일 {count}개 남음</div>
    </TodoHeadBlock>
  );
}

export default TodoHead;

import React from "react";
import styled from "styled-components";

const SNSButton = styled.button`
  width: 100%;
  padding: 10px;
  background-color: ${(props) => props.bgColor || "#ced4da"};
  color: white;
  font-size: 16px;
  font-weight: bold;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 10px;

  &:hover {
    opacity: 0.9;
  }
`;

const SNSButtons = () => {
  return (
    <>
      <SNSButton bgColor="#3b5998">페이스북</SNSButton>
      <SNSButton bgColor="#F7FE2E" style={{ color: "black" }}>
        카카오
      </SNSButton>
      <SNSButton bgColor="#DB4437">구글</SNSButton>
    </>
  );
};

export default SNSButtons;

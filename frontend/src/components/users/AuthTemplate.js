import React from "react";
import { createGlobalStyle } from "styled-components";
import styled from "styled-components";
import SNSButtons from "./SnsButtons";

const GlobalStyle = createGlobalStyle`
  body {
    background: #e9ecef;
  }
`;

const AuthTemplateBlock = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: #e9ecef;
`;

const AuthBox = styled.div`
  width: 360px;
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
`;

const Title = styled.h1`
  font-size: 24px;
  color: #343a40;
  margin-bottom: 30px;
`;

const AuthTemplate = ({ title, children, snsButtons }) => {
  return (
    <>
      <GlobalStyle />
      <AuthTemplateBlock>
        <AuthBox>
          <Title>{title}</Title>
          {children}
        </AuthBox>
      </AuthTemplateBlock>
    </>
  );
};

export default AuthTemplate;

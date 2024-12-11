import React, { useState } from "react";
import AuthTemplate from "./AuthTemplate";
import { Link, useNavigate } from "react-router-dom";
import styled from "styled-components";
import SNSButtons from "./SnsButtons";
import { loginUsers } from "../../api/UsersApi";

const Form = styled.form`
  display: flex;
  flex-direction: column;
`;

const Input = styled.input`
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ced4da;
  border-radius: 4px;
`;

const Button = styled.button`
  padding: 10px;
  background-color: #38d9a9;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;

  &:hover {
    background-color: #32c79a;
  }
`;

const Divider = styled.div`
  border-top: 1px solid #ced4da;
  margin: 20px 0;
`;

const SnsTitle = styled.p`
  font-size: 14px;
  color: #868e96;
  margin-bottom: 10px;
`;

const LoginForm = ({ setIsLoggedIn }) => {
  const sessionStorage = window.sessionStorage;
  const [error, setError] = useState(null);
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const handleInputChange = (formName, value) => {
    setFormData({
      ...formData,
      [formName]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const responseData = await loginUsers(formData);
      const { token, userId } = responseData;
      sessionStorage.setItem("accessToken", token);
      sessionStorage.setItem("userId", userId);
      // 성공 시 리디렉션 처리 등
      setIsLoggedIn(true);
      navigate("/todos");
    } catch (error) {
      alert(error.message); // 서버에서 전달된 에러 메시지를 alert로 표시
    }
  };

  return (
    <AuthTemplate title="MINI TodoList">
      <Form onSubmit={handleSubmit}>
        <Input
          type="email"
          placeholder="Email"
          value={formData.email}
          onChange={(e) => {
            handleInputChange("email", e.target.value);
          }}
        />
        <Input
          type="password"
          placeholder="Password"
          value={formData.password}
          onChange={(e) => {
            handleInputChange("password", e.target.value);
          }}
        />
        <Button type="submit">Login</Button>
        <div style={{ marginTop: "10px", textAlign: "right" }}>
          <Link to="/signup" style={{ color: "#38d9a9", fontSize: "14px" }}>
            회원가입
          </Link>
        </div>
        <div style={{ marginTop: "20px" }}>
          <Divider />
          <SnsTitle>SNS 로그인</SnsTitle>
          <SNSButtons />
        </div>
      </Form>
    </AuthTemplate>
  );
};

export default LoginForm;

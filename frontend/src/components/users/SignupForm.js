import React, { useState } from "react";
import AuthTemplate from "./AuthTemplate";
import { Link, useNavigate } from "react-router-dom";
import styled from "styled-components";
import { saveUsers } from "../../api/UsersApi";

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

const SignupForm = ({ onSignup }) => {
  const navigate = useNavigate();
  const [error, setError] = useState(null);

  const [formData, setFormData] = useState({
    email: "",
    name: "",
    password: "",
    passwordConfirm: "",
  });

  const handleInputChange = (formName, value) => {
    setFormData({
      ...formData,
      [formName]: value,
    });
  };

  const handleSignup = async (e) => {
    e.preventDefault();
    if (formData.password !== formData.passwordConfirm) {
      alert("비밀번호가 일치하지 않습니다.");
      return;
    }

    try {
      const response = await saveUsers(formData);
      console.log(response);
      if (response.status === 201) {
        alert("회원가입 완료!");
        navigate("/login");
      } else {
        setError("회원가입 실패. 다시 시도해주세요.");
      }
    } catch (error) {
      setError("회원가입 중 오류가 발생했습니다.");
      console.error("Error saving user:", error);
    }
  };

  return (
    <AuthTemplate title="회원가입">
      <Form onSubmit={handleSignup}>
        <Input
          type="email"
          placeholder="이메일"
          value={formData.email}
          onChange={(e) => {
            handleInputChange("email", e.target.value);
          }}
        />
        <Input
          type="text"
          placeholder="이름"
          value={formData.name}
          onChange={(e) => {
            handleInputChange("name", e.target.value);
          }}
        />
        <Input
          type="password"
          placeholder="비밀번호"
          value={formData.password}
          onChange={(e) => {
            handleInputChange("password", e.target.value);
          }}
        />
        <Input
          type="password"
          placeholder="비밀번호 확인"
          value={formData.passwordConfirm}
          onChange={(e) => {
            handleInputChange("passwordConfirm", e.target.value);
          }}
        />
        <Button type="submit">회원가입</Button>
        <div style={{ marginTop: "10px", textAlign: "center" }}>
          <Link to="/login" style={{ color: "#38d9a9", fontSize: "14px" }}>
            로그인 페이지로 이동
          </Link>
        </div>
      </Form>
    </AuthTemplate>
  );
};

export default SignupForm;

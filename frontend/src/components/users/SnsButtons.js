import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";

const FacebookButton = styled.button`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 360px;
  height: 41px;
  background-color: #1877f2; /* Facebook Blue */
  border: none;
  border-radius: 5px;
  color: white;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s;
  margin-bottom: 10px;

  &:hover {
    background-color: #145dbf; /* Darker Facebook Blue */
  }

  &:active {
    background-color: #0d4a94; /* Even darker */
  }

  .icon {
    margin-right: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  img {
    width: 20px;
    height: 20px;
  }
`;

const KakaoButton = styled.button`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 360px;
  height: 41px;
  background-color: #fee500; /* Kakao Yellow */
  border: none;
  border-radius: 5px;
  color: #3c1e1e; /* Kakao Text Color */
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s;
  margin-bottom: 10px;

  &:hover {
    background-color: #fddc00; /* Slightly darker Kakao Yellow */
  }

  &:active {
    background-color: #fccc00; /* Even darker */
  }

  .icon {
    margin-right: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  img {
    width: 20px;
    height: 20px;
  }
`;

const GoogleButton = styled.button`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 360px;
  height: 41px;
  background-color: white; /* 구글 버튼은 기본적으로 흰색 */
  border: 1px solid #dadce0;
  border-radius: 5px;
  color: #3c4043; /* 구글 텍스트 색상 */
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s, box-shadow 0.3s;

  &:hover {
    background-color: #f8f9fa; /* 살짝 회색 */
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.15);
  }

  &:active {
    background-color: #eceff1;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  }

  .icon {
    margin-right: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  img {
    width: 20px;
    height: 20px;
  }
`;

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
  const handleKakaoLogin = () => {
    axios.get("/auth/kakao").then((res) => {
      const restApiKey = res.data.client_id;
      const redirectUri = res.data.redirect_uri;
      const responseType = res.data.response_type;
      // 카카오 로그인 페이지로 리다이렉트
      const kakaoUrl = `https://kauth.kakao.com/oauth/authorize?client_id=${restApiKey}&redirect_uri=${redirectUri}&response_type=${responseType}`;
      window.location.href = kakaoUrl; // 페이지 리다이렉트
    });
  };

  return (
    <>
      <FacebookButton>
        <div className="icon">
          <img
            src="https://upload.wikimedia.org/wikipedia/commons/5/51/Facebook_f_logo_%282019%29.svg"
            alt="Facebook logo"
          />
        </div>
        Log in with Facebook
      </FacebookButton>
      <KakaoButton onClick={handleKakaoLogin}>
        <div className="icon">
          <img
            src="https://developers.kakao.com/assets/img/about/logos/kakaolink/kakaolink_btn_small.png"
            alt="Kakao logo"
          />
        </div>
        Log in with Kakao
      </KakaoButton>
      <GoogleButton>
        <div className="icon">
          <img
            // 이미지 URL을 변경한 부분
            src="https://apis.google.com/js/platform.js?onload=renderButton"
            alt="Google logo"
          />
        </div>
        Sign in with Google
      </GoogleButton>
    </>
  );
};

export default SNSButtons;

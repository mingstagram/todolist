import axios from "axios";
import React, { useEffect, useState } from "react";

const KakaoRedirection = () => {
  const code = new URL(window.location.href).searchParams.get("code");
  const [hasRedirected, setHasRedirected] = useState(false); // 페이지가 리다이렉트된 여부를 추적

  useEffect(() => {
    // `code`가 존재하고, 리다이렉트가 한 번만 발생하도록 추적
    if (code && !hasRedirected) {
      setHasRedirected(true); // 리다이렉트된 상태로 설정

      axios
        .get(`/auth/kakaoRegist?code=${code}`)
        .then((response) => {
          // 성공 시 처리
          console.log("Response:", response.data);
        })
        .catch((error) => {
          // 실패 시 처리
          console.error("Error:", error);
        });
    }
  }, [code, hasRedirected]); // `code`와 `hasRedirected`를 의존성 배열로 설정

  return <div>Redirecting...</div>;
};

export default KakaoRedirection;

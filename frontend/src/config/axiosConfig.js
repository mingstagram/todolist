import axios from "axios";

// Axios 인스턴스 생성
const axiosInstance = axios.create({
  baseURL: "http://localhost:8001", // API의 기본 URL 설정
  headers: {
    "Content-Type": "application/json", // 기본 헤더 설정
  },
});

// 요청 인터셉터를 통해 Authorization 헤더 추가
axiosInstance.interceptors.request.use(
  (config) => {
    const token = sessionStorage.getItem("accessToken"); // sessionStorage에서 토큰 가져오기
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`; // Authorization 헤더에 토큰 추가
    }
    return config;
  },
  (error) => {
    return Promise.reject(error); // 에러 처리
  }
);

export default axiosInstance; // axios 인스턴스 반환

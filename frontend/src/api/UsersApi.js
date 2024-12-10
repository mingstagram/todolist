import axios from "axios";

export const saveUsers = async (userData) => {
  try {
    const response = await axios.post("/users", userData);
    return response;
  } catch (error) {
    console.error("Error creating users:", error);
    throw error;
  }
};

export const loginUsers = async (userData) => {
  try {
    const response = await axios.post("/users/login", userData);

    if (response.data.code === "0000") {
      return response.data.data; // 성공 시 데이터 반환
    } else {
      throw new Error(response); // 서버에서 전달한 메시지 사용
    }
  } catch (error) {
    if (error.response) {
      // 서버 에러 응답 처리
      throw new Error(error.response.data.message || "Login failed");
    } else {
      // 네트워크 오류 처리
      throw new Error("Network error");
    }
  }
};

# TodoList 프로젝트

**간단한 TodoList 웹 애플리케이션**

이 프로젝트는 사용자가 할 일을 추가, 수정, 삭제할 수 있는 기능을 제공하는 미니 프로젝트입니다. Go와 React를 사용하여 개발되었으며, MySQL을 데이터베이스로 사용합니다.

## 주요 기능

- 할 일 추가
- 할 일 목록 보기
- 할 일 수정
- 할 일 삭제

---

## 설치 방법

1. **저장소 클론**

   ```bash
   git clone https://github.com/mingstagram/todolist.git
   cd todolist
   ```

2. **백엔드 실행**

- 필요한 의존성 설치
  ```bash
  go mod tidy
  ```
- 서버 실행
  ```bash
  go run cmd/server/main.go
  ```

3. **프론트엔드 실행**

- 프론트엔드 디렉토리로 이동 후 의존성 설치
  ```bash
  cd frontend
   npm install
   npm start
  ```

---

## 사용 방법

1. 웹 브라우저에서 http://localhost:3000으로 접속
2. 할 일을 추가하거나 삭제하며 테스트

---

## 기술 스택

- 프론트엔드: React
- 백엔드: Go (Gorilla/Mux)
- 데이터베이스: MySQL
- 배포: Docker, AWS EC2

---

## 기여 방법

1. 이 저장소를 포크합니다.
2. 새로운 브랜치를 생성합니다
   ```bash
   git checkout -b feature/my-feature
   ```
3. 변경 사항을 커밋합니다.
   ```bash
   git commit -m "Add my feature"
   ```
4. 브랜치를 푸시합니다.
   ```bash
   git push origin feature/my-feature
   ```
5. 풀 리퀘스트를 생성합니다.

---
 

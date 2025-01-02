# 1. 빌드 환경 설정
# Node.js 환경에서 React 앱 빌드
FROM node:18 AS builder
WORKDIR /app

# package.json과 package-lock.json을 먼저 복사하고, 종속성 설치
COPY package.json package-lock.json ./
RUN npm install

# 애플리케이션 코드 복사 및 빌드
COPY . .
RUN npm run build

# 2. 실행 환경 설정
# Nginx를 사용해 정적 파일 제공
FROM nginx:1.23
COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

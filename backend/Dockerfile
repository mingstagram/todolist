# Golang을 이용해 빌드를 위한 이미지
FROM golang:1.20 AS builder

# 작업 디렉토리 설정 (Go 애플리케이션의 루트 디렉토리로 설정)
WORKDIR /app

# 의존성 파일 복사
COPY go.mod go.sum ./

# Go 의존성 다운로드
RUN go mod download

# 애플리케이션 소스 코드 복사 (전체 백엔드 소스 코드 복사)
COPY . .   

# 애플리케이션 빌드
RUN go build -o main ./cmd/server   

# 실행 환경 이미지 (Debian 베이스)
FROM debian:bookworm

# 작업 디렉토리 설정
WORKDIR /app

# 빌드된 바이너리 파일 복사
COPY --from=builder /app/main .

# 실행 명령어 설정
CMD ["./main"]

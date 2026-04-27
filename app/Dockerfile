# Stage 1: Build ตัวแอปพลิเคชัน
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o aradinboard main.go

# Stage 2: นำตัวแอปที่ Build เสร็จแล้วมารันใน Image เปล่าๆ (เพื่อลดขนาด Image)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/aradinboard .
EXPOSE 8080
CMD ["./aradinboard"]
FROM golang:1.24.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apt-get update && \
    apt-get install -y libgl1-mesa-dev xorg-dev && \
    rm -rf /var/lib/apt/lists/*
RUN go build -o job-visualizer ./cmd/app

FROM debian:bookworm-slim
WORKDIR /app
RUN apt-get update && \
    apt-get install -y libgl1-mesa-glx && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/JobData.xlsx .
COPY --from=builder /app/job-visualizer .

CMD ["./job-visualizer", "--headless"]
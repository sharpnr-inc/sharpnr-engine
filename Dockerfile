# ---------- Build stage ----------
FROM golang:1.24.4 AS builder

WORKDIR /src

# cache deps
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .
COPY .env.production .env.production
# build binary (name changed to avoid folder conflict)
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ---------- Runtime stage ----------
FROM ubuntu:22.04

WORKDIR /app

# 🔥 Install CA certificates
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# 👇 COPY FROM CORRECT LOCATION
COPY --from=builder /src/server .
COPY --from=builder /src/sharp_mailer_engine/templates ./sharp_mailer_engine/templates
# ⭐ ADD THIS (IMPORTANT)
COPY --from=builder /src/.env.production .env.production

EXPOSE 8080

CMD ["./server"]

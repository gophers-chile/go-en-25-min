# Stage 1: Build de aplicacion
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Stage 2: Runtime de aplicacion
FROM alpine:3.19
WORKDIR /app

# Copia el binario desde el stage de build
COPY --from=builder /app/main .

# Expone el puerto y define el comando por defecto
EXPOSE 8080
CMD ["./main"]
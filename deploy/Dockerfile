###########################################
# Construcción (Desarrollo)
###########################################
FROM golang:1.23.6-alpine AS dev

ENV GOLANG_VERSION 1.23.6

RUN go install github.com/air-verse/air@latest

ENV PATH "/go/bin:${PATH}"

WORKDIR /app

# Copiar dependencias primero
COPY go.mod ./
RUN go mod download && go mod verify

# Copiar el código fuente, plantillas y archivos generados
COPY . .

EXPOSE 3000

###########################################
# Construcción (Producción)
###########################################
FROM golang:1.23.6-alpine AS build

ENV GOLANG_VERSION 1.23.5

WORKDIR /go/src/kybermed_server

# Copiar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente y las plantillas generadas
COPY . .

# Construir el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -installsuffix cgo -o /go/src/kybermed_server/app ./cmd/kybermed_server

###########################################
# Etapa 4: Release
###########################################
FROM alpine:latest AS release

WORKDIR /app

# Copiar el binario construido
COPY --from=build /go/src/kybermed_server/app .

RUN apk -U upgrade \
  && apk add --no-cache dumb-init ca-certificates \
  && chmod +x /app/app

EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/app"]

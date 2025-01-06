###########################################
# Etapa 1: Generación de Código con Templ
###########################################
FROM ghcr.io/a-h/templ:latest AS generate

WORKDIR /app
# Copiar solo las plantillas necesarias
COPY web/template /app/template
RUN ["templ", "generate"]

###########################################
# Etapa 2: Desarrollo (Air)
###########################################
FROM golang:1.23-alpine AS dev

RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copiar dependencias primero
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente, plantillas y archivos generados
COPY . .
COPY --from=generate /app/template /app/web/template

EXPOSE 3000
CMD ["air", "-c", ".air.toml"]

###########################################
# Etapa 3: Construcción (Producción)
###########################################
FROM golang:1.23-alpine AS build

WORKDIR /go/src/kybermed

# Copiar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente y las plantillas generadas
COPY . .
COPY --from=generate /app/template /go/src/kybermed/web/template

# Construir el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -installsuffix cgo -o /go/src/kybermed/app ./cmd/kybermed

###########################################
# Etapa 4: Release
###########################################
FROM alpine:latest AS release

WORKDIR /app

# Copiar el binario construido
COPY --from=build /go/src/kybermed/app .

RUN apk -U upgrade \
  && apk add --no-cache dumb-init ca-certificates \
  && chmod +x /app/app

EXPOSE 3000
ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/app"]

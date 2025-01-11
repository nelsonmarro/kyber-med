FROM golang:1.23-alpine AS dev

RUN go install github.com/a-h/templ/cmd/templ@latest \
  && go install github.com/air-verse/air@latest

ENV PATH="/go/bin:${PATH}"

WORKDIR /app

# Copiar dependencias primero
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente, plantillas y archivos generados
COPY . .

EXPOSE 3000
CMD ["air", "-c", ".air.toml"]

###########################################
# Etapa 3: Construcción (Producción)
###########################################
FROM golang:1.23-alpine AS build

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

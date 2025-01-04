FROM golang:1.23-alpine AS dev

RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copia los archivos de dependencias primero (para hacer cache de mod)
COPY go.mod go.sum ./
RUN go mod download

# Copia todo el proyecto
COPY . .

# Expone el puerto que uses en desarrollo
EXPOSE 3000

# Comando principal en desarrollo
CMD ["air", "-c", ".air.toml"]


###########################################
# Etapa 2: Build (para producir binario)
###########################################
FROM golang:1.23-alpine AS build

WORKDIR /go/src/kybermed

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -installsuffix cgo -o /go/src/kybermed/app ./cmd/kybermed


FROM alpine:latest AS release
WORKDIR /app

# Copiamos binario compilado
COPY --from=build /go/src/kybermed/app .

# (Opcional) Copia carpeta static si la usas
RUN mkdir -p static
COPY --from=build /go/src/kybermed/static ./static

RUN apk -U upgrade \
  && apk --update add git make bash build-base \
  && apk add --no-cache dumb-init ca-certificates \
  && chmod +x /app/app

EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/app"]

FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN go test ./... && CGO_ENABLED=0 go build -o /out/server ./cmd/server

FROM alpine:3.20
RUN addgroup -S app && adduser -S app -G app && mkdir -p /data && chown app:app /data
WORKDIR /app
COPY --from=build /out/server /app/server
USER app
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/server"]

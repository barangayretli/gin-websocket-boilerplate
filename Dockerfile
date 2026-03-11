FROM golang:1.23-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app .

FROM alpine:3.19
RUN adduser -D -u 1000 appuser
USER appuser
COPY --from=build /app /app
EXPOSE 8000
ENTRYPOINT ["/app"]

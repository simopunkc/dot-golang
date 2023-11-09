##
## Build
##
FROM golang:1.21.1-alpine AS build

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o blog cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=build /app/env.example /app/.env
COPY --from=build /app/blog /app/blog

USER nonroot:nonroot

CMD ["./blog"]
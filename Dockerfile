FROM golang:1.19-alpine as build

WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:latest as production
COPY --from=build /api/openhms-api /app/bin/openhms

CMD ["/app/bin/openhms"]
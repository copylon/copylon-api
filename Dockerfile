FROM golang:1.17 as build

WORKDIR /api
COPY . .
RUN go mod tidy && go mod verify
RUN go build

FROM alpine:latest as production
RUN apk update && apk upgrade && apk add --no-cache libc6-compat
COPY --from=build /api/openhms /app/bin/openhms

CMD ["/app/bin/openhms"]
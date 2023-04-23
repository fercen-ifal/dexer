FROM golang:1.20-alpine as build

WORKDIR /dexer/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /dexer/build

FROM alpine

WORKDIR /app
COPY --from=build ./dexer/build ./dexer/.env ./

RUN chmod +x /app/build

EXPOSE 8080

CMD [ "/app/build" ]

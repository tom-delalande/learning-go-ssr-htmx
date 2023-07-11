FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY server/*.go ./server/
COPY server/view/*.html ./server/view/
COPY server/view/styles/*.css ./server/view/styles/
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app ./server

CMD ["./app"]

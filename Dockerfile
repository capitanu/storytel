FROM golang:alpine AS BUILDER

WORKDIR /go/src/github.com/capitanu/storytel

COPY go.mod .
COPY go.sum . 
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o main .

FROM scratch AS DEPLOYER
COPY --from=BUILDER /go/src/github.com/capitanu/storytel/main .

EXPOSE 8180
CMD ["./main"]

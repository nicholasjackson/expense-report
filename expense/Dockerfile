FROM golang:1.17.1 as builder

RUN mkdir -p /go/src/github.com/nicholasjackson/expense-report/expense

COPY . /go/src/github.com/nicholasjackson/expense-report/expense

WORKDIR /go/src/github.com/nicholasjackson/expense-report/expense

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o expense-report


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /go/src/github.com/nicholasjackson/expense-report/expense/expense-report /app/

ENTRYPOINT [ "/app/expense-report" ]
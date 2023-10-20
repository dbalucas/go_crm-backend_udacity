FROM golang:1.21.3 AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .
RUN go mod download

RUN  go build -o /app/crm-backend

FROM scratch

WORKDIR /app
COPY /static /app/static
COPY --from=builder /app/crm-backend /app/crm-backend

EXPOSE 3000

CMD [ "/app/crm-backend" ]
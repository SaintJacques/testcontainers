FROM golang:1.18-alpine as build

WORKDIR /testcntrns
ADD . /testcntrns

RUN go build -o /tmp/testcntrns ./*.go

FROM alpine as run 

COPY --from=build /tmp/testcntrns /bin/testcntrns
COPY --from=build /testcntrns/migrations/ ./migrations/

RUN chmod +x /bin/testcntrns

CMD ["/bin/testcntrns"]
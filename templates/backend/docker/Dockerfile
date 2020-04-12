FROM alpine:latest

COPY dist/{{AppName}} .
RUN mkdir -p /log/logs

CMD ["./{{AppName}}"]


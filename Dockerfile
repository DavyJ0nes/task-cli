# Multistage Build

### CREATE DOCKERMASTER USER
FROM alpine:3.6 AS alpine
RUN adduser -D -u 10001 dockmaster

## MAIN IMAGE
FROM scratch
LABEL Name=task
LABEL Author=davyj0nes

COPY --from=alpine /etc/passwd /etc/passwd

ADD app /
USER dockmaster

ENTRYPOINT ["./task-cli"]

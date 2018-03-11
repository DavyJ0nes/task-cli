# Multistage Build

## CREATE DOCKERMASTER USER
FROM alpine:3.6 AS alpine
RUN adduser -D -u 10001 dockmaster
RUN touch /root/empty

## MAIN IMAGE
FROM scratch
LABEL Name=task
LABEL Author=davyj0nes

COPY --from=alpine /etc/passwd /etc/passwd

ADD task_static /app/task
# USER dockmaster
WORKDIR /app
ENV HOME=/app

COPY --from=alpine /root/empty /app/.tasks/empty

ENTRYPOINT ["./task"]

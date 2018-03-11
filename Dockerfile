# Multistage Build

## CREATE DOCKERMASTER USER
FROM alpine:3.6 AS alpine
RUN touch /tmp/empty

## MAIN IMAGE
FROM scratch
LABEL Name=task
LABEL Author=davyj0nes

ADD task_static /app/task
COPY --from=alpine /tmp/empty /app/.tasks/empty

WORKDIR /app
ENV HOME=/app

ENTRYPOINT ["./task"]

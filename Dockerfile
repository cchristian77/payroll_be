FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY /env.json /app/env.json

COPY payroll-api-build /app

CMD [ "/app/payroll-api-build" ]
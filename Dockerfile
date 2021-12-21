FROM golang:1.13-buster

WORKDIR /app

COPY /out/helloWorld /app/

EXPOSE 80

CMD [ "/app/helloWorld" ]
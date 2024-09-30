FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update
RUN apt-get install -y build-essential python3 python3-pip openjdk-11-jdk
RUN useradd -m runner
USER runner

WORKDIR /home/runner

COPY main.c .

CMD ["sh", "-c", "gcc main.c -o main && ./main"]

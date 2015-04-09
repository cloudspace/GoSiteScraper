FROM ubuntu:latest
RUN apt-get install -y libxml2-dev ca-certificates
ADD ./scraper /scraper
ADD ./microservice.yml /microservice.yml

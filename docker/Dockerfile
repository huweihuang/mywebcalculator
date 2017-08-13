FROM docker.io/centos
MAINTAINER huweihuang
COPY mywebcalculator /mywebcalculator/
COPY views /mywebcalculator/views/
COPY docker/run.sh /mywebcalculator/
EXPOSE 8080
WORKDIR /mywebcalculator
CMD ["/mywebcalculator/run.sh"]

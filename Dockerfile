FROM ubuntu:latest

COPY ./Veeee /opt/
RUN chmod +x /opt/Veeee
EXPOSE 8080
CMD ["/opt/Veeee"]
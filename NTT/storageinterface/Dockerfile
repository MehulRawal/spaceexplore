FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=test_database
ENV MYSQL_USER=test_user
ENV MYSQL_PASSWORD=secret

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306
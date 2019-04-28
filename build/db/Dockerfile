FROM mysql:5.7.25
ENV MYSQL_DATABASE=dreba MYSQL_ROOT_PASSWORD=dreba
COPY init.sql /docker-entrypoint-initdb.d/
COPY conf.d/* /etc/mysql/conf.d/
EXPOSE 3306

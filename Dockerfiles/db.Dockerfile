FROM postgres

COPY . /docker-entrypoint-initdb.d

EXPOSE 5432
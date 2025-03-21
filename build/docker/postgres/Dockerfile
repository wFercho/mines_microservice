FROM postgres:15

# Copia el script de inicialización a la carpeta de PostgreSQL
COPY init.sql /docker-entrypoint-initdb.d/

# Configura el usuario y base de datos por defecto
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=mines_db
# Use an official PostgreSQL runtime as a parent image
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydb

# Optionally, expose the PostgreSQL port (5432) if needed
# EXPOSE 5432

# Copy an initialization SQL script into the container
COPY init.sql /docker-entrypoint-initdb.d/

# The init.sql file should contain SQL commands to initialize your database
# For example, you can create tables and insert data
# Example init.sql content:
# CREATE TABLE mytable (id serial PRIMARY KEY, name VARCHAR(255));
# INSERT INTO mytable (name) VALUES ('John');

# You can also mount volumes, manage settings, and configure authentication as needed

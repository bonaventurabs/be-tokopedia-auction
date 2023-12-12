# Use the official PostgreSQL image from the Docker Hub
FROM postgres:alpine

# Environment variables to configure PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=12345
ENV POSTGRES_DB=tokopedia-auction

# Expose the PostgreSQL default port
EXPOSE 5432

# Set up a volume for PostgreSQL data
VOLUME /var/lib/postgresql/data
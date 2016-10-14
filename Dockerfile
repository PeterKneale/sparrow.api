FROM golang

# Create the directory where the application will reside
RUN mkdir /app
RUN mkdir /app/swagger

# Copy the application files (needed for production)
ADD sparrow.api /app/sparrow.api
ADD swagger/* /app/swagger/

# Set the working directory to the app directory
WORKDIR /app

# Expose the application on port 8080.
EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT /app/sparrow.api
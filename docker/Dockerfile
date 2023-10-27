# Download Image docker
FROM golang:1.21

# Set destination for COPY
WORKDIR /app/build

# Download Go modules
COPY go.* .
RUN go mod download

# Copy the source code
COPY ./src ./src

# Copy scripts sql
COPY ./sql ../sql

# Copy env file
COPY .env ../

# Build
RUN cd ./src && CGO_ENABLED=0 GOOS=linux go build -o ../product-api

# Arguments available
ARG USER_DATABASE=root \
    PASSWORD_DATABASE=root \
    DATABASE_NAME=productdb \
    DATABASE_HOST \
    DATABASE_PORT=3306 \

# Environments 
ENV USER_DATABASE=$USER_DATABASE \
    PASSWORD_DATABASE=$PASSWORD_DATABASE \
    DATABASE_NAME=$DATABASE_NAME \
    DATABASE_HOST=$DATABASE_HOST \
    DATABASE_PORT=$DATABASE_PORT \
    SERVER_PORT=$SERVER_PORT

# Expose port to access the application
EXPOSE 5000

# Run
CMD ["./product-api"]
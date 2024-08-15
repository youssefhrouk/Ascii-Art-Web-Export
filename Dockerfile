# Use the official Alpine base image
FROM golang:1.22.3-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Install bash in the container
RUN apk add bash

# we copy everything in the root directory to /app directory
COPY . .

# Build the app with optional configuration 
RUN go build -o /ascii

# tells Docker that the comtainer listens on specified network ports at runtime 
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/ascii" ]

# Set the authors of the image
LABEL authors="asoudri, yhrouk"

# Add a description for the image
LABEL description="This Docker image is for the Ascii-Art-Web project, which is a web server written in Go that converts text into ASCII art."

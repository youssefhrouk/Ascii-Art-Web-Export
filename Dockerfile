# Use the official Alpine base image
FROM golang:1.22.3-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# we copy everything in the root directory to /app directory
COPY . .

# Build the app with optional configuration 
RUN go build -o /3ab9or

LABEL version="1.0"
LABEL description="ASCII-ART-WEB"

# tells Docker that the comtainer listens on specified network ports at runtime 
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/3ab9or" ]
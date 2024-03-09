# Use the official Go image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the application
RUN go build -o main .

# Expose port 80
EXPOSE 80

# Run the application
CMD ["./main"]

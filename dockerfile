# Use official Node.js image as the base
FROM node:18

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json first (if available)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Expose port for the application (adjust if needed)
EXPOSE 3000

# Command to start the application
CMD ["node", "server/server.js"]


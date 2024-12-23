
先整理下我们有哪些 Behavior 写了的，然后把现有的 Behavior 整理下，再看看还能加哪些新的 Behavior

# LoadGenerator (for Train-Ticket): A New Version loadgenerator for Microservice Systems
The Train-Ticket LoadGenerator is a dedicated tool designed to simulate traffic for the Train Ticket Booking System, which is based on a microservice architecture containing 41 microservices. This tool is primarily developed in Go, leveraging its performance and simplicity to effectively test and benchmark the system.

## Behavior Logic Graph
![Behavior-Logic-Graph0.png](assest/images/Screenshot 2024-12-23 at 1.17.14 AM.png)
![Behavior-Logic-Graph1.png](assest/images/Screenshot 2024-12-23 at 1.16.10 AM.png)
### Existing Behavior:


## Environment Setup and Deployment Guide
This guide provides the necessary steps to set up the environment and deploy the application.
#### Prerequisites
To get started, ensure you have the following installed and configured:
1. **Goland IDE**: Recommended for development with Go.
2. **Go Modules**: Run the following command to tidy up dependencies:
   ```bash
   go mod tidy
   ```
#### Deployment and Running the Application
To deploy and run the application, follow these steps:
1. Set the `BASE_URL` environment variable(Replace 'http://10.10.10.220:30222' with the corresponding address):
   ```powershell
   $env:BASE_URL = "http://10.10.10.220:30222"
   ```
2. Start the application:
   ```bash
   go run main.go
   ```
That's it! The application should now be running and accessible at the specified `BASE_URL`.


---

For any issues or further details, feel free to check the documentation or raise an issue in the repo :D

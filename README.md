# test-data-population-app
An app to generate data for a testing environment that will help me practice unfamiliar concepts later on.

This is a simple Go application that connects to a Redis service in a Kubernetes cluster, populates it with test data, and provides an HTTP endpoint to retrieve the stored messages.

## Prerequisites

Make sure you have the following installed:

- [Minikube](https://minikube.sigs.k8s.io/docs/start/)

## Getting Started

1. **Start Minikube:**

   ```bash
   minikube start
   ```

2. **Build Docker Image:**

   ```bash
   docker build -t your-docker-username/test-data-population-app:latest .
   ```

3. **Push Docker Image:**

   ```bash
   docker push your-docker-username/test-data-population-app:latest
   ```

4. **Apply Kubernetes Resources:**

   ```bash
   kubectl apply -f deployment.yaml
   kubectl apply -f service.yaml
   ```

5. **Port Forwarding:**

   Open a new terminal and run:

   ```bash
   kubectl port-forward service/test-data-population-app 8080:80
   ```

6. **Access the Application:**

   Open Postman and make a GET request to:

   ```http
   http://localhost:8080/messages
   ```

   Or you can use your web browser or any other HTTP client.

7. **Stop Minikube:**

   When you're done testing, you can stop Minikube:

   ```bash
   minikube stop
   ```

## Notes

- Make sure that your Docker Hub account is configured, and you have the necessary permissions to push the Docker image.
- The Kubernetes resources (`deployment.yaml` and `service.yaml`) assume a basic Redis service running in the `go-api-practice` namespace. Adjust them if your setup is different.
- Replace `your-docker-username` with your actual Docker Hub username.

Feel free to customize the Go application, Kubernetes resources, and this README according to your specific requirements.


This `README.md` provides step-by-step instructions on how to set up and run your Go application on Minikube. Ensure to replace placeholders like `your-docker-username` with your actual Docker Hub username.
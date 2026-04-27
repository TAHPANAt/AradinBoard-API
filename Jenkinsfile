pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps { echo '1. Pull Code from GitHub' }
        }
        stage('Build & Test') {
            steps { sh 'cd app && go build -o aradinboard main.go' }
        }
        stage('Docker Build & Push') {
            steps { echo '2. Build and Push to Docker Hub' }
        }
        stage('Deploy (IaC)') {
            steps { echo '3. Run Terraform & Ansible' }
        }
    }
}
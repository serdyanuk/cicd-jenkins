pipeline {
    agent any
    stages {
        stage('Compile') {
            steps {
                sh "go version"
                sh "go build -v ./..."
            }
        }
        stage('Test') {
            steps {
                sh "go test -v ./..."
            }
        }
    }
}
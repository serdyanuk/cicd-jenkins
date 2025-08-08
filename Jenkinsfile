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
        stage('Code Coverage') {
          steps {
             sh '''
                  go test -coverprofile=coverage.out ./...
                  $(go env GOPATH)/bin/gocover-cobertura < coverage.out > coverage.xml
                '''
          }
        }
        stage('Publish Coverage') {
          steps {
            recordCoverage tools: [[parser: 'COBERTURA', pattern: 'coverage.xml']]
          }
        }
    }

  post {
    always {
      archiveArtifacts artifacts: 'coverage.*', fingerprint: true
    }
  }
}
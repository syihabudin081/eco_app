pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'eco_app'
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/syihabudin081/eco_app.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app main.go'
            }
        }

        stage('Docker Build') {
            steps {
                sh "docker build -t ${DOCKER_IMAGE} ."
            }
        }

        stage('Docker Run') {
            steps {
                // Use the 'withCredentials' block to access credentials
                withCredentials([string(credentialsId: 'your-env-variable-id', variable: 'YOUR_ENV_VAR')]) {
                    sh "docker run --env YOUR_ENV_VAR=${YOUR_ENV_VAR} -d ${DOCKER_IMAGE}"
                }
            }
        }

        stage('Cleanup') {
            steps {
                sh 'docker system prune -f'
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: '**/app', fingerprint: true
            cleanWs()
        }
    }
}

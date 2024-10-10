pipeline {
  agent any
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
        withCredentials(bindings: [string(credentialsId: 'your-env-variable-id', variable: 'YOUR_ENV_VAR')]) {
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
  environment {
    DOCKER_IMAGE = 'eco_app'
    DATABASE_URL = '"postgresql://eco_db_owner:Hbe04QnvdoLa@ep-little-morning-a173f8ef.ap-southeast-1.aws.neon.tech/eco_db?sslmode=require"'
  }
  post {
    always {
      archiveArtifacts(artifacts: '**/app', fingerprint: true)
      cleanWs()
    }

  }
}
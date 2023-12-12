pipeline {
  agent {
    kubernetes {
	inheritFrom 'dind'
    }
  }
    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Pushing like a pregnant woman...'
                container('dind') {
		   script {
		    docker.withRegistry("https://registry.service.consul", "docker-registry-personal") {
                    def customImage = docker.build("dorero:${env.BUILD_ID}")
                    customImage.push()
                    }
		   }
                }
            }
        }
    }
}

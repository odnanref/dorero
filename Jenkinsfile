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
                    def customImage = docker.build("registry.service.consul/dorero", 'docker-registry-personal')
                    customImage.push()
                    }
                }
            }
        }
    }
}

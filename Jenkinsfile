pipeline {
    agent any

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
                echo 'Deploying....'
                script {
                    def customImage = docker.build("registry.service.consul:5000/dorero")
                    customImage.push()
                }
            }
        }
    }
}

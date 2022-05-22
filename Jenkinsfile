pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                
                script {
                    def customImage = docker.build("registry.service.consul:5000/dorero")
                    customImage.push()
                }
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
            }
        }
    }
}
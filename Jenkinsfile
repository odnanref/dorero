pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                
                curl https://get.docker.com/ > dockerinstall && chmod 777 dockerinstall && ./dockerinstall
                
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
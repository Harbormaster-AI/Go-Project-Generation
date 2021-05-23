pipeline {
    agent { docker { image 'Jenkinsfile_Docker_App_Image()' } }
    stages {
        stage('install') {
            steps {

                sh go get github.com/gorilla/mux 
                sh go get -u gorm.io/gorm
                sh go get -u gorm.io/driver/mysql
                sh go get github.com/joho/godotenv
            }
            
        stage('build') {
            steps {
                sh cd src/go-demo
                sh go mod init go-demo
                sh go mod tidy
                sh go build
                sh go test go-demo/api/test
            }
        }
    }
}

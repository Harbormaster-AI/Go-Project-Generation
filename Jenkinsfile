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
                sh cd src/godemo
                sh go mod init godemo
                sh go mod tidy
                sh go build
                sh go test godemo/api/test
            }
        }
    }
}

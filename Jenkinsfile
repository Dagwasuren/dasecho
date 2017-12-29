pipeline {
    agent any

        stages {
            stage('Build') {
                steps {
                    echo 'Building..'
                    sh "docker build -t dasecho ."

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
                    sh "rsync -avzP ./* /data/www/dasecho.net/"                
                    sh "cd /data/www/web-svc && ls -al && ./force-replace.sh dasecho &&./reload.sh"
                }
            }
        }
}



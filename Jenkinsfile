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
					sh "cd /data/www/dasecho.net && git pull --rebase && cd /data/www/web-svc && docker-compose create --build --force-recreate dasecho "
				}
			}
		}
}

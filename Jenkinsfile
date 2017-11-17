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
					sh "rsync -avzP ./* /data/www/dasecho.net/ && cd /data/www/web-svc && docker-compose create --build --force-recreate dasecho && docker-compose restart dasecho"
					sh "docker run --rm -v $(pwd):/www -v /data:/data gobuffalo/buffalo:v0.10.1 sh -c "cd /www;buffalo db migrate"
				}
			}
		}
}

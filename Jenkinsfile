pipeline {
    agent any


    parameters{
        string(name: 'tag', defaultValue: 'dev', description: 'Image Tag')
        booleanParam(name: 'deploy', defaultValue: true, description: 'Deploy in machine')
        booleanParam(name: 'dockerHub', defaultValue: false, description: 'Push to Docker Hub')
    }
    stages {
        stage('Create Tag') {
            steps {
                sh "chmod 777 ./scripts/create_tag.sh"
                sh "./scripts/create_tag.sh"
            }
        }
        stage('SCM UI Author') {
            steps {
                dir('ds-ui-author') {
                  git branch: "$BRANCH_NAME", url: 'https://github.com/appveen/ds-ui-author.git'
                }
            }
        }
        stage('SCM UI Appcenter') {
            steps {
                dir('ds-ui-appcenter') {
                  git branch: "$BRANCH_NAME", url: 'https://github.com/appveen/ds-ui-appcenter.git'
                }
            }
        }
        stage('SCM UI Swagger') {
            steps {
                dir('ds-ui-swagger') {
                  git branch: "$BRANCH_NAME", url: 'https://github.com/appveen/ds-ui-swagger.git'
                }
            }
        }
        stage('Build') {
            steps {
                sh "chmod 777 ./scripts/build.sh"
                sh "./scripts/build.sh"
            }
        }
        stage('Prepare YAML') {
            steps {
                sh "chmod 777 ./scripts/prepare_yaml.sh"
                sh "./scripts/prepare_yaml.sh"
            }
        }
        stage('Push to ECR') {
            steps {
                sh "chmod 777 ./scripts/push_ecr.sh"
                sh "./scripts/push_ecr.sh"
            }
        }
        stage('Save to S3') {
            when {
                expression {
                    params.dockerHub  == true
                }
            }
            steps {
                sh "chmod 777 ./scripts/push_s3.sh"
                sh "./scripts/push_s3.sh"
            }
        }
        stage('Deploy') {
            when {
                expression {
                    params.deploy == true
                }
            }
            steps {
                sh "chmod 777 ./scripts/deploy.sh"
                sh "./scripts/deploy.sh"
            }
        }
        stage('Push to Docker Hub') {
            when {
                expression {
                    params.dockerHub  == true
                }
            }
            steps {
                sh "chmod 777 ./scripts/push_hub.sh"
                sh "./scripts/push_hub.sh"
            }
        }
    }
}
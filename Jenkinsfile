pipeline {
    agent any

    stages {
        stage ('Clone') {
            steps {
                checkout([$class: 'GitSCM',
                    branches: [[name: '*/master']],
                    doGenerateSubmoduleConfigurations: false,
                    extensions: [], submoduleCfg: [],
                    userRemoteConfigs: [[url: 'https://github.com/devops-ws/learn-pipeline-go']]]
                )
            }
        }

        stage ('Build Code') {
            agent {
                label 'golang'
            }
            steps {
                container ('golang') {
                    sh label: 'Build Go Source Code', script: 'make build'
                }
            }
        }

        stage ('Build Image') {
            agent {
                label 'docker'
            }
            steps {
                container ('docker') {
                    sh label: 'Build Image', script: 'make image'
                }
            }
        }

        stage ('Charts Package') {
            agent {
                label 'helm'
            }
            steps {
                container ('helm') {
                    sh label: 'Charts Package', script: 'helm package go-server-charts'
                }
            }
        }

        stage ('Charts Upload') {
            agent {
                kubernetes {
                    yaml '''
                        apiVersion: v1
                        kind: Pod
                        metadata:
                        labels:
                            app: helm-push
                        spec:
                        containers:
                        - name: helm
                            image: surenpi/helm-push:v0.0.1
                            command:
                            - cat
                            tty: true
                    '''
                    label 'helm-push'
                }
            }
            steps {
                container ('helm') {
                    sh label: 'Charts Upload', script: '''
                    helm push mychart/ http://localhost:8080
                    '''

                    /**
                    # Or you can use the curl command directly
                    sh label: 'Charts Upload', script: '''
                    curl -L --data-binary "@go-server-0.1.0.tgz" http://localhost:8080/api/charts
                    '''
                    */
                }
            }
        }
    }
}
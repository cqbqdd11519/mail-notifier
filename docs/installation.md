# Installation Guide

## Prerequisites
- SMTP Server, ID/PW for the server

## Components
- **Secret** specifying SMTP server user/password
- Mail-sender server **Deployment/Service**

## Procedure
1. Configure Secret with SMTP server account
    ```bash
    SMTP_SERVER=<SMTP Server Address>
    SMTP_USER=<SMTP User ID>
    SMTP_PW=<SMTP User PW>
    NAMESPACE=<Namespace the server to be deployed>
    
    curl https://raw.githubusercontent.com/cqbqdd11519/mail-notifier/master/deploy/secret.yaml.template -s | \
    sed "s/<SMTP Address (IP:PORT)>/${SMTP_SERVER}/g" | \
    sed "s/<SMTP User ID>/${SMTP_USER}/g" | \
    sed "s/<SMTP User PW>/${SMTP_PW}/g" | \
    kubectl apply --namespace ${NAMESPACE} -f -
    ```
2. Deploy Service/Deployments for mail-sender server
    ```bash
    kubectl apply --namespace ${NAMESPACE} --filename https://raw.githubusercontent.com/cqbqdd11519/mail-notifier/master/deploy/service.yaml
    kubectl apply --namespace ${NAMESPACE} --filename https://raw.githubusercontent.com/cqbqdd11519/mail-notifier/master/deploy/server.yaml
    ```

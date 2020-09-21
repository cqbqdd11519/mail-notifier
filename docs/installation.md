# Installation Guide

## Prerequisites
- External SMTP Server (e.g., Gmail...), ID/PW for the server  
(We don't provide internal SMTP server due to security issues)

## Components
- **Secret** specifying SMTP server user/password
- Mail-sender server **Deployment/Service**

## Images Required
* mail-sender-server ([tmaxcloudck/mail-sender-server:v0.0.4](https://hub.docker.com/layers/tmaxcloudck/mail-sender-server/v0.0.4/images/sha256-3d87f419d056132690bd7cdcb5aab1abe0021ae12b4efd50a8b7c0be7a44dd86?context=explore))
* mail-sender-client ([tmaxcloudck/mail-sender-client:v0.0.4](https://hub.docker.com/layers/tmaxcloudck/mail-sender-client/v0.0.4/images/sha256-0364005e432a67e839cee04cdb0ebb5d925eb4427fd248f346566300f890d046?context=explore))

## Procedure
1. Configure Secret with SMTP server account
    ```bash
    SMTP_SERVER=<SMTP Server Address>
    SMTP_USER=<SMTP User ID>
    SMTP_PW=<SMTP User PW>
    NAMESPACE=<Namespace the server to be deployed>
    
    curl https://raw.githubusercontent.com/cqbqdd11519/mail-notifier/master/deploy/secret.yaml.template -s | \
    sed "s/<SMTP Address (IP:PORT)>/'${SMTP_SERVER}'/g" | \
    sed "s/<SMTP User ID>/'${SMTP_USER}'/g" | \
    sed "s/<SMTP User PW>/'${SMTP_PW}'/g" | \
    kubectl apply --namespace ${NAMESPACE} -f -
    ```
2. Deploy Service/Deployments for mail-sender server
    ```bash
    kubectl apply --namespace ${NAMESPACE} --filename https://raw.githubusercontent.com/cqbqdd11519/mail-notifier/master/deploy/service.yaml
    kubectl apply --namespace ${NAMESPACE} --filename https://raw.githubusercontent.com/cqbqdd11519/mail-notifier/master/deploy/server.yaml
    ```

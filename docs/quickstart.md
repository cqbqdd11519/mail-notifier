# Quick Start Guide

1. Configure ConfigMap containing mail receiver ([example/configmap.yaml](../example/configmap.yaml))
   * Each line of `data.users` filed should be in form of \<user name\>=\<email address\>
    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: approver-dev
      namespace: default
    data:
      users: |
        shkim-tmax.co.kr=sunghyun_kim3@tmax.co.kr
    ```
2. Execute Pod containing mail-sender-client container (using Job, in this example)
    ```yaml
    apiVersion: batch/v1
    kind: Job
    metadata:
      name: send-mail
    spec:
      template:
        spec:
          containers:
            - name: client
              image: tmaxcloudck/mail-sender-client:latest
              env:
              - name: MAIL_SERVER
                value: http://mail-sender.<NAMESPACE of mail-sender-server>:9999/
              - name: MAIL_FROM
                value: <email sender address>
              - name: MAIL_SUBJECT
                value: <email title>
              - name: MAIL_CONTENT
                value: |
                  <email content>
              volumeMounts:
                - name: mail-list
                  mountPath: /tmp/config # mountPath cannot be changed
          restartPolicy: Never
          volumes:
            - name: mail-list
              configMap:
                name: approver-dev # Should be the name of pre-configured configmap
    ```

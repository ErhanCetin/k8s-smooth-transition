# News Tracker Application 
---
 ![fe-page-narrow](https://user-images.githubusercontent.com/10308201/99090312-bd0a6d80-25ce-11eb-9c95-a18f8f6093d2.png)"
---
- This project's been created to be used in k8s articles, microservices as well in the feature, to give details to readers. I want to a reader to make some practices with my articles. You can easily apply the YAML file after doing prerequisites following the below section to check the project quickly. Basically, there are a couple of services that aim to fetch tech news from [newsapi] to show in the web app. Please check the components section to get insight into the project.
- If you’re using this sample application, please ★Star this repository to show your interest!

## Service Architecture
---
 ![application-architecture](https://user-images.githubusercontent.com/10308201/98815665-cd341880-2427-11eb-8800-f5054711997c.png)


## Components / Applications
---
##### 1. NewsTrackerProducer Service

- It is implemented with **Python** to fetch the tech news from [newsapi] and write them down to activemq queue.

##### 2. ActiveMQ Service

- This service is used to persist the data coming from [newsapi] in the queue. --> used by NewsTrackerProducer and NewsTrackerConsumer Services.

##### 3. NewsTrackerConsumer Service

- It is implemented with **golang** to fetch the news from Activemq service and write them down to mongo db.
   
##### 4. Mongodb Service

- It is used to persist **news** . --> used by Mongo-Express Service, NewsTrackerConsumer and NewsTrackerProducer Services.

##### 5. Mongo-Express Service

- It is used to check **Mongo DB** via user interface. 
   
##### 6. NewsTrackerApi Service 

- It is implemented with **Spring Boot and SpringData MongoTemplate**. The service serves the **news** data. --> used by NewsTrackerFE Service  
   
##### 7. NewsTrackerFE Service 

- It is implemented with **[next.js]**. This service presents all news list which is coming from NewsTrackerApi service to the end user.
 
### Steps between the services
 1. Fetching news from [newsapi] and write them to activemq. --> by NewsTrackerProducer Service  
 2. Consuming news from the activemq and put them into the mongo db. --> by NewsTrackerConsumer Service
 3. Fetching news data from the Mongo Db and expose them for the FE. --> by NewsTrackerApi Service
 4. Fetching the news from the NewsTrackerApi to show them to the end user in formatted web page. --> by NewsTrackerFE Service   
 
 
### Installation for running Kubernetes locally
---
* **Prerequisites:** 
   * Install git : https://git-scm.com/downloads
   * Install docker: https://docs.docker.com/get-docker/
   * Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
   * Install minikube: https://kubernetes.io/docs/tasks/tools/install-minikube/ 
   * An API Key of [newsapi] is used to fetch the news. It is used by the News Producer Service. I already put my API key to the configuration. But this key is limited to 100 requests per day. I recommend you to get your own API key from [newsapi-free] : 
     * Select the "Developer" option , After registration , Go your account and copy API Key.
        ![newspi org-account](https://user-images.githubusercontent.com/10308201/99146187-1e394c00-2676-11eb-82e8-3320923c62a6.png)
     * Change YAML file, I hope you already clone the project to your local.
       * Find : /k8s-smooth-transition/k8s/apps/04-newsproducer.yaml and change NEWSAPI_AUTHORIZATION with your API Key.
         * ![newsapi-org-config-key](https://user-images.githubusercontent.com/10308201/99146276-dff05c80-2676-11eb-9dac-ebd062438a82.png) 
   * Don't forget 
     * Check the version compatibility of docker,kubectl and minikube
     * Increase **resources** of docker in your local (Recommended resources : 4CPU & 5GB) .
     
* **The whole project was tested in a local env with :**
    *   Kubernetes v1.18.3
    *   minikube v1.12.1 on Darwin 10.15.4
    *   Docker 19.03.2 ( Resources 4CPU & 5GB)
* **Execute commands to verify installation :** 
    ```sh
      $ docker --version
      $ docker-compose --version        
      $ docker-machine --version        
      $ minikube version                
      $ kubectl version 
    ```  
                 
 
* **Start minikube**
    ```sh
      $ minikube start 
      $ minikube status 
    ```  

* **Install all services to Kubernetes instance on minikube:**
    - Clone the project to your local :
        ```sh
        $ git clone https://github.com/ErhanCetin/k8s-smooth-transition.git
        ``` 
    - Check the minikube :
         ```sh
          $ minikube status 
        ``` 
    - Open a terminal and go to k8s-smooth-transition/k8s/apps, you will see [all-news-tracker-services]   
    - Execute kubectl command to run all services up .
        ```sh
        $ kubectl apply -f .
        ```
**---- >  Installation is done.**
### What should you have after running all-services up in Kubernetes ?
---
- After installation steps , you should see resouces below in your local .
    - **Pods Resources**
         ```
            $ kubectl get pods # ignore generated number in pods names. 
        ```
        | NAME                                         | READY | STATUS | RESTARTS | AGE | 
        | -------------------------------------------- | ----- | ------ | -------- | --- |
        | activemq-deployment-5b8ff65958-mhdtr           |1/1   |Running|     1|    16m
        | mongodb-stateful-0                             | 1/1  |Running|     1|    16m
        | newsapi-deployment-b678b9b7f-f8v8q             | 1/1  |Running|     0|    16m
        | newsconsumer-deployment-7fb89b858b-lrttz       | 1/1  |Running|     0|    16m
        | newsfe-deployment-77756f7bd4-wnxfc             | 1/1  |Running|     0|    16m
        | newsmongdbexpress-deployment-589d48b946-fwv92  | 1/1  |Running|     0|    24s
        | newsproducer-cronjob-1605045660-sg227          | 0/1  |Completed|   0|    66s
       
    - **Deployment Resources**    
         ```
            $ kubectl get deployment 
        ```
        |NAME                           |READY   |UP-TO-DATE   |AVAILABLE   |AGE
        |--|--|--|--|--
        |activemq-deployment            |1/1     |1            |1           |41m
        |newsapi-deployment             |1/1     |1            |1           |41m
        |newsconsumer-deployment        |1/1     |1            |1           |41m
        |newsfe-deployment              |1/1     |1            |1           |41m
        |newsmongdbexpress-deployment   |1/1     |1            |1           |41m
    
    - **Service Resources**
        ```
            $ kubectl get svc
        ```
        |NAME                     |TYPE        |CLUSTER-IP       |EXTERNAL-IP   |PORT(S)              |AGE
        | ------------------------|------------------- | ----- | ------ | -------- | --- |
        |activemq-service         |ClusterIP   |10.101.216.209   |<none>        |8161/TCP,61613/TCP   |25m
        |kubernetes               |ClusterIP   |10.96.0.1        |<none>        |443/TCP              |28m
        |mongodb-service          |ClusterIP   |None             |<none>        |27017/TCP            |25m
        |mongodbexpress-service   |ClusterIP   |10.103.19.72     |<none>        |8081/TCP             |25m
        |newsapi-service          |ClusterIP   |10.111.245.72    |<none>        |8080/TCP             |25m
        |newsfe-service           |NodePort    |10.100.4.7       |<none>        |3000:30144/TCP       |25m

    - **ConfigMap Resources**
        ```
            $ kubectl get svc
        ```
        |NAME                       |DATA   |AGE
        | --- | ---|---|
        |activemq-configmap         |3      |30m
        |mongodb-configmap          |1      |30m
        |mongodbexpress-configmap   |2      |30m
        |newsapi-configmap          |3      |30m
        |newsconsumer-configmap     |10     |30m
        |newsfe-configmap           |3      |30m
        |newsproducer-configmap     |11     |30m
    
    - **Secret Resources**
        ```
            $ kubectl get secret
        ```
        |NAME                             |TYPE                                  |DATA   |AGE
        |--|---|--|--|
        |activemq-admin-password          |Opaque                                |1      |34m
        |newsproducer-activemq-password   |Opaque                                |1      |34m 
    - **Cronjob Resource**
        ```
            $ kubectl get cronjob
        ```
        |NAME                   |SCHEDULE      |SUSPEND   |ACTIVE   |LAST SCHEDULE   |AGE
        |--|--|--|--|--|--|
        |newsproducer-cronjob   |0/1 * * * *   |False     |1        |14s             |38m
    
    - **StatefulSet Resource** 
        ```
            $ kubectl get statefulset
        ```
        |NAME               |READY   |AGE
        |--|--|--|
        |mongodb-stateful   |1/1     |40m

    - **Persistentvolume Resource** 
        ```
            $ kubectl get persistentvolume
        ```
        
        |NAME                                       |CAPACITY   |ACCESS MODES   |RECLAIM POLICY   |STATUS      |CLAIM                    |STORAGECLASS   |REASON   |AGE
        |--|--|--|--|--|--|--|--|--|
        |activemq-pv                                |100M       |RWO            |Retain           |Available                                                    |43m
        |fee-pv                                     |100M       |RWO            |Retain           |Available                                                    |43m
        |pvc-e897f978-c911-4d6c-816b-dc608caf6bbb   |100M       |RWO            |Delete           |Bound       |default/activemq-claim   |standard              |  43m
        |pvc-f9be7acb-883d-4518-950b-b505eaf8c080   |100M       |RWO            |Delete           |Bound       |default/fee-claim        |standard              |  43m

    - **PersistentVolumeClaim Resource** 
        ```
            $ kubectl get persistentvolume
        ```
        |NAME             |STATUS   |VOLUME                                     |CAPACITY   |ACCESS MODES   |STORAGECLASS   |AGE
        |--|--|--|--|--|--|--
        |activemq-claim   |Bound    |pvc-e897f978-c911-4d6c-816b-dc608caf6bbb   |100M       |RWO            |standard       |44m
        |fee-claim        |Bound    |pvc-f9be7acb-883d-4518-950b-b505eaf8c080   |100M       |RWO            |standard       |44m


##### ➜ **Finally everthing is ready to vist our website to read the tech news :** 
- Command below will open a browser for the tech news. Enjoy it.  
       ```
        $ minikube service newsfe-service
      ```
- check the browser , you should see the page below
- ![fe-page-narrow](https://user-images.githubusercontent.com/10308201/99090312-bd0a6d80-25ce-11eb-9c95-a18f8f6093d2.png)
---
### License : MIT
---
---
#### Contact : [![Linkedin](https://i.stack.imgur.com/gVE0j.png) erhancetin](https://www.linkedin.com/in/erhancetin/)
&nbsp;

[//]:# (Reference Links)
[newsapi]: <https://newsapi.org>
[newsapi-free]: <https://newsapi.org/pricing>
[next.js]: <https://nextjs.org>
[all-news-tracker-services]: <https://github.com/ErhanCetin/k8s-smooth-transition/tree/develop/k8s/apps> 
[fe-page-wide]: https://user-images.githubusercontent.com/10308201/99090289-b67bf600-25ce-11eb-8562-3f008a543ec3.png


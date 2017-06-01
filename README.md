<p align="center">
  <img src="https://cloud.githubusercontent.com/assets/4391815/26681386/05b857c4-46ab-11e7-8c71-15a46d886834.png">
</p>


<!--
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Turbonomic

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# **KubeTurbo** 

----

1. [Overview](#overview)
2. [Getting Started](#getting-started)
3. [Kubeturbo Use Cases](#use-cases)
4. [Coming Soon](#coming-soon)

----

## Overview 

Kubeturbo leverages [Turbonomic](https://turbonomic.com/) patented analysis engine to provide visibility and control across the entire stack, to assure the performance of running micro-service in Kubernetes Pods, as well as the efficiency of underlying infrastructure.

## Getting Started 
### Prerequisite
* Turbonomic 5.9+ installation
* Kubernetes 1.4+
### Kubeturbo Installation
* [Deploy Kubeturbo](https://github.com/turbonomic/kubeturbo/tree/master/deploy)

![screen shot 2017-06-01 at 9 50 06 am 2](https://cloud.githubusercontent.com/assets/4391815/26682910/106ebb7c-46b0-11e7-88dd-9d0d073c0620.png)

## Use Cases
* Full-Stack Visibility by leveraging 50+ existed Turbonomic controllers, from on-prem DataCenter to major public cloud providers. No more shadow IT
  * From Load Balancer all the way down to your physical Infrastructure
  * Real-Time resource monitoring across entire DataCenter
  * Real-Time Cost visibility for your public cloud deployment
![screen shot 2017-06-01 at 9 49 48 am 2](https://cloud.githubusercontent.com/assets/4391815/26682938/2cb28296-46b0-11e7-892d-8f06c02e7072.png)
![screen shot 2017-06-01 at 9 58 43 am](https://cloud.githubusercontent.com/assets/4391815/26683161/eab420ba-46b0-11e7-9fb2-4faad6f08025.png)

* Provide Rescheduler capability by leveraging Turbonomic analysis engine (Require Kubeturbo to be the scheduler)
  * Consolidating Pod in real time to increase node efficiency
  * Reschedule Pod in advance to prevent suffering resource congestion from the underlying node
  *	Reschedule Pod to new node added to the cluster
  *	Reschedule Pods that peak together to different nodes, to avoid performance dropping
![screen shot 2017-06-01 at 9 50 00 am 2](https://cloud.githubusercontent.com/assets/4391815/26682996/59ccfdce-46b0-11e7-9b6d-6ccad13df3fa.png)


* Right-Sizing your Pod and your entire IT stack
  *	Combining Turbonomic real time performance monitor and real time analysis engine, Turbonomic is able to provide the right-sizing information for every individual pod and also the entire IT stack.
  * Right-sizing up your Pod limit, if necessary, to avoid OOM
  * Right-sizing down your Pod requested resource, if necessary, to avoid resource overprovisioning or overspending in public cloud deployment.
![screen shot 2017-06-01 at 9 56 55 am](https://cloud.githubusercontent.com/assets/4391815/26683094/b330c350-46b0-11e7-91e5-a1db65a89d50.png)


## Coming Soon
* Cluster Federation Control Plane
  * Complete visibility for your K8s deployments across different underlying infrastructures
  * Create affinity/anti-affinity policy directly from Turbonomic UI
  * Improve cost efficiency by consolidating workload across deployments and identifying cheapest region and provider to deploy your workload
* What-If Planer 
  * A complete What-If sandbox to help you plan your IT changes in advance
  * Plan for workload change: Add/Remove Containers
  * Plan for infrastructure change: Add/Remove/Replace hardware
  * Migrate your Kubernetes deployment from On-Prem to the cloud, how much is the cost
  * Cluster Consolidation for federated clusters



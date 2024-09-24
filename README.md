# Advanced Software Design project
### Team name: Tech Tornados
### Project title: Load Balancer with Service Turn-Off
### Team Members:
- Dmitriy Kara
- Daniil Mikulik
- Ekaterina Karavaeva
- Nikita Dumkin

## Summary

The "Load Balancer with Service Turn-Off" project aims to enhance the functionality of a vanilla Kubernetes (k8s) distribution by enabling it to scale applications down to zero instances when they are not in use. This is particularly relevant for machine learning (ML) applications that often have unique computational requirements and can be slow to start. The service will use advanced, event-driven automata to manage application scaling in response to real-time monitoring data, optimizing resource usage and ensuring that applications can scale up quickly when needed.

## Stakeholders

1. DevOps Engineers

    Responsible for managing Kubernetes clusters and ensuring high availability and scalability of applications.

2. ML Engineers and Data Scientists

    Require efficient and scalable infrastructure to deploy and manage model training and inference tasks.

3. Service Reliability Engineers

    Oversee the underlying infrastructure and ensure optimal resource allocation and utilization, observe core service availability and performance metrics.

4. Cloud Service Providers

    Provide the necessary infrastructure and want to offer cutting-edge PaaS system for load balancing.

5. End Users

    Benefit from improved performance, resource utilization and reduced latency of the applications they use.

## Expected Needs

1. Cost Efficiency

    Ability to reduce operational costs by scaling unused application instances to zero.

2. Performance

    Ensuring that applications scale up quickly from zero to meet demand without significant latency.

3. Resource Optimization

    Better utilization of underlying computational resources, particularly for ML applications.

4. Ease of Integration

    Seamlessly integrate with existing Kubernetes environments without requiring significant configuration changes.

5. Security

    Ensure that the scaling service does not introduce vulnerabilities to the deployment.


## Features

1. Zero Scaling

    Allow applications to scale down to zero instances when not in use and automatically scale up when required.

2. Event-Driven Scaling

    Use complex event processing to trigger scaling actions based on defined monitoring metrics and thresholds.

3. Custom Resource Allocation

    Tailor the allocation of compute resources specifically for ML applications, ensuring they get the necessary resources when scaling up.

4. Integration with Monitoring Tools

    Seamlessly integrate with existing monitoring tools to gather real-time data for making scaling decisions.

5. Dashboards and Alerts

    Provide user-friendly dashboards to visualize events and performance metrics, and set up alerts for critical events.

6. Automata-Based Policy Engine

    Deploy a sophisticated policy engine employing complex automata for precise and customizable scaling rules.

7. API for Customization

    Offer an API that allows for customization and tuning of the scaling algorithms to fit specific application needs.

## Constraints

1. Latency

    Ensuring minimal latency when scaling from zero to active instances, especially for ML applications which can be resource-intensive.

2. Compatibility

    Must be compatible with a wide range of Kubernetes distributions and configurations, without requiring extensive modifications.

3. Scalability

    The solution must manage high levels of concurrency and large clusters efficiently.

4. Security

    Must adhere to security best practices and compliance requirements to avoid introducing vulnerabilities.

5. Resource Limits

    Operate within the constraints of available computational resources and quotas set by end users.

6. Reliability

    Ensure high availability and fault tolerance within the scaling service itself to prevent outages or scaling issues.

7. Complexity

    The service should strive to reduce operational complexity, ensuring it remains accessible to users without significant management overhead.


# Task 3. Presentation
https://docs.google.com/presentation/d/1IaeXMkLVIqLq7L1JCghJjV8ALhiMkqwj1jUcEOtdsyg/edit?usp=sharing

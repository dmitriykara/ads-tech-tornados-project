# Use case text explanation

1. Edit Balancer Configuration
   - Actor: Engineer (DevOps / ML engineer)
   - Description: This use case allows the engineer to modify the load balancer settings to apply scaling, healthcheck or allocation policies.

2. Edit Healthcheck Policy
   - Actor: Engineer (DevOps / ML engineer)
   - Description: The engineer configures the healthcheck policy to change the instance availability is monitoring rules.

3. Check Instances Availability
   - Actor: Application (MLOps framework)
   - Description: Regular checks are conducted to verify that required instances of the application are running and have a valid configuration.

4. Edit Scaling Policy
   - Actor: Engineer (DevOps / ML engineer)
   - Description: The engineer adjusts the scaling policy to manage resource allocation efficiently, including the ability to autoscale down to zero when needed.

5. Autoscale Down to Zero
   - Actor: Application (MLOps framework)
   - Description: Automatically scale down to zero resources when they are not needed to optimize cost and efficiency.

6. Edit Allocation Policy
   - Actor: Engineer (DevOps / ML engineer)
   - Description: Allows the engineer to modify how resources are allocated based on current application needs and policies. Also required for resources constraints' management.

7. Allocate Resources
   - Actor: Application (MLOps framework)
   - Description: Allocates computing resources dynamically based on the current workload and policy settings.

8. View Performance Dashboards
   - Actor: Engineer (DevOps / ML engineer)
   - Description: The engineer can view dashboards that illustrate key performance and resource utilization metrics.

9. Aggregate Observability Metrics
   - Actor: Monitoring System
   - Description: Collects and aggregates various metrics to provide a comprehensive view of the system's performance, availability and resource utilization. Metrics aggregation is required for future visualization.

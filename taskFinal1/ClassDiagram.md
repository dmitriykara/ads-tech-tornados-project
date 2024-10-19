# Class diagram description

[Link to class diagram](https://drive.google.com/file/d/1uZnKI4C8L5o2sd2wmaLDSjnipEQ8TWfj/view?usp=sharing)

The UML class diagram depicts a comprehensive design for a project focusing on load balancing with service turn-off features. Here’s a detailed description of the components and their interactions.

**N.B.** It is important, that all key components of the system (first of all `PolicyController`) use kubernetes `ConfigMap` format for specification configuration.

## Key Components:

1. Node
   - Attributes:
     - id: String
     - title: String
     - pods: Pod[]
   - Methods:
     - create(): Creates a new pod.
     - read(podId: String): Reads pod information.
     - update(pod: Pod): Updates a pod.
     - delete(podId: String): Deletes a pod.
     - applyPodSpec(podIds: String[]): Applies specifications to pods.
   - Represents an logical object for node in the Kubernetes cluster, containing information about pods.

2. NodesController (Controller)
   - Attributes:
     - id: String
     - nodes: Node[]
   - Methods:
     - checkResourceLimit(): Checks overall resource limits.
     - reportNodeStatus(nodeId: String): Reports the status of the node.
     - handlePolicyUpdate(): Reacts on ploicy updates, such as AllocationPolicy or HealthcheckPolicy

3. Pod
   - Attributes:
     - id: String
     - containerIds: String[]
     - resources: ConfigMap
     - status: String
     - spec: ConfigMap
   - Methods:
     - start(): Starts the pod.
     - stop(): Stops the pod.
     - destroy(): Destroys the pod.
     - restart(): Restarts the pod.
     - applyPodSpec(): Applies specifications to the pod.
     - checkResourceLimit(): Checks pod resource limits.

4. EntityController (Interface)
   - Methods:
     - create(): Creates an entity.
     - read(entityId: String): Reads an entity.
     - update(entity: Entity): Updates an entity.
     - delete(entityId: String): Deletes an entity.

5. PolicyController (Controller)
   - Attributes:
     - id: String
     - policies: Policy[]
   - Methods:
     - create(): Creates a policy.
     - read(policyId: String): Reads a policy.
     - update(policy: Policy): Updates a policy.
     - delete(policyId: String): Deletes a policy.
     - applyId: String

6. Policy (Abstract)
   - Attributes:
     - id: String
     - spec: ConfigMap
   - Methods:
     - apply(): Emits a signal of policy change, handled by other parts of a system.

7. AllocationPolicy
   - Extends: Policy
   - Methods:
     - apply(): Updates policy spec and emits signal for Kubelet to reallocate resources.
     - checkResourceLimit(): Checks resource limits before policy spec update.
     - allocateResources(): Reserves resources for future allocations.

8. ScalingPolicy
   - Extends: Policy
   - Methods:
     - apply(): Updates policy spec and emits signal for Balancer to update balancing algorithm or traffic distribution rules.

9. HealthCheckPolicy
    - Extends: Policy
    - Methods:
      - apply(): Updates policy spec and emits signal for Kubelet to check pods liveness and readiness.

10. Balancer
    - Attributes:
      - id: String
      - algorithmType: String
      - trafficDistribution: String
    - Methods:
      - updateAlgorithm(): Updates stored balancing algorithm type.
      - applyAlgorithm(): Applies chosen algorithm type to runtime specification.
      - emitMonitoringEvent(): Emits monitoring events about balancing system (like RPS, resource utilization, availability status).

11. ObservabilitySystem
    - Attributes:
      - id: String
      - monitoringEventIds: String[]
      - alertIds: String[]
      - logs: String[]
    - Methods:
      - writeLog(): Writes balancer logs.
      - sendAlerts(alertIds: String[]): Sends alerts based on received events from balancer.
      - storeMonitoringEvent(): Receives monitoring events from balancer and stores them.

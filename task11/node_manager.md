# Node Manager MongoDB Schema Documentation

## Collections

### 1. `node_manager`
Manages the nodes within the Kubernetes environment, including their status, resources, and associated pods.

#### Schema
```json
{
  "_id": "ObjectId",
  "nodeId": "string", // Unique identifier for the node
  "name": "string", // Node name
  "status": "string", // Node status: "Active", "Inactive", "Error"
  "resources": {
    "cpu": "number", // Total CPU resources in cores
    "memory": "number" // Total memory resources in MB
  },
  "pods": [
    {
      "podId": "string", // Unique identifier for the pod
      "name": "string", // Pod name
      "status": "string", // Pod status: "Running", "Pending", "Failed"
      "containers": [
        {
          "containerId": "string", // Unique identifier for the container
          "name": "string", // Container name
          "status": "string" // Container status: "Running", "Stopped", "Error"
        }
      ]
    }
  ],
  "location": {
    "region": "string", // Geographical region of the node
    "zone": "string" // Zone within the region
  },
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```


## Indexes

- **Single index** on `nodeId`:
  ```javascript
  db.node_manager.createIndex({ "nodeId": 1 });
  ```
  Ensures fast lookup of nodes by their unique ID.

- **Compound index** on `status` and `resources.cpu`:
  ```javascript
  db.node_manager.createIndex({ "status": 1, "resources.cpu": -1 });
  ```
  Optimizes queries for filtering nodes by status and CPU resources.

- **Geospatial index** on `location.region` and `location.zone`:
  ```javascript
  db.node_manager.createIndex({ "location.region": 1, "location.zone": 1 });
  ```
  Supports efficient queries for region and zone-specific nodes.


## Key Selection and Query Patterns

### Key Selections:
- **Primary Key (_id)**: Automatically indexed for unique identification.
- **Indexes**:
  - Single index on `nodeId`.
  - Compound index on `status` and `resources.cpu`.
  - Geospatial index on `location.region` and `location.zone`.

### Query Patterns:
1. **Retrieve a specific node by its unique ID**:
   ```javascript
   db.node_manager.findOne({ "nodeId": "<NODE_ID>" });
   ```

2. **List all active nodes sorted by CPU usage (descending)**:
   ```javascript
   db.node_manager.find({ "status": "Active" }).sort({ "resources.cpu": -1 });
   ```

3. **Fetch nodes in a specific region and zone**:
   ```javascript
   db.node_manager.find({ "location.region": "us-east", "location.zone": "zone-1" });
   ```

4. **Get all pods running on a specific node**:
   ```javascript
   db.node_manager.aggregate([
       { $match: { "nodeId": "<NODE_ID>" } },
       { $project: { "pods": 1 } }
   ]);
   ```

5. **Retrieve all nodes with failed pods**:
   ```javascript
   db.node_manager.find({ "pods.status": "Failed" });
   ```

6. **Search for containers with errors within all pods**:
   ```javascript
   db.node_manager.find({ "pods.containers.status": "Error" });
   ```

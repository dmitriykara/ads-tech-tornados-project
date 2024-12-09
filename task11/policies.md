# Policies MongoDB Schema Documentation

## Collections

### 1. `policies`
Stores various types of policies, such as Allocation, Scaling, and HealthCheck.

#### Schema
```json
{
  "_id": "ObjectId",
  "id": "string", // Unique identifier for each policy
  "type": "string", // Policy type: "Allocation", "Scaling", "HealthCheck"
  "name": "string", // Human-readable policy name
  "spec": {
    "config": [
      // Configuration details as key-value pairs
    ],
    "createdBy": "string", // User or system that created the policy
    "createdAt": "timestamp",
    "updatedAt": "timestamp",
    "description": "string", // Brief description of the policy
    "status": "string" // "Active", "Inactive", "Deprecated"
  },
  "rules": [
    {
      "ruleId": "string",
      "condition": "string", // e.g., "CPU > 80%"
      "action": "string" // e.g., "Scale Up"
    }
  ]
}
```

#### Indexes
- **Single index** on `id`:
  ```javascript
  db.policies.createIndex({ "id": 1 });
  ```
  Ensures quick retrieval of policies by their unique ID.
- **Compound index** on `type` and `status`:
  ```javascript
  db.policies.createIndex({ "type": 1, "spec.status": 1 });
  ```
  Optimizes fetching policies by type and status.
- **Text index** on `name` and `spec.description`:
  ```javascript
  db.policies.createIndex({ "name": "text", "spec.description": "text" });
  ```
  Supports full-text search for policy names and descriptions.


## Key Selection and Query Patterns

### Key Selections:
- **Primary Key (_id)**: Automatically indexed for unique identification.
- **Indexes**:
  - Single index on `id`.
  - Compound index on `type` and `spec.status`.
  - Text index on `name` and `spec.description`.

### Query Patterns:
1. **Retrieve a specific policy by its unique ID**:
   ```javascript
   db.policies.findOne({ "id": "<POLICY_ID>" });
   ```

2. **List all active Scaling policies**:
   ```javascript
   db.policies.find({ "type": "Scaling", "spec.status": "Active" });
   ```

3. **Search policies by name or description containing a keyword**:
   ```javascript
   db.policies.find({ $text: { $search: "scaling" } });
   ```

4. **Retrieve all policies created by a specific user**:
   ```javascript
   db.policies.find({ "spec.createdBy": "admin" });
   ```

5. **Fetch policies with a specific rule condition**:
   ```javascript
   db.policies.find({ "rules.condition": "CPU > 80%" });
   ```
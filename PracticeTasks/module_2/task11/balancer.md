# Balancer MongoDB Schema Documentation

## Collections

### 1. `balancer_events`
Stores events related to balancing operations.

#### Schema
```json
{
  "_id": "ObjectId",
  "eventType": "string",
  "details": "object",
  "status": "string", // e.g., "pending", "processed"
  "createdAt": "timestamp",
  "updatedAt": "timestamp",
  "priority": "integer" // Optional field for priority sorting
}
```

#### Indexes
- **Compound index** on `status` and `priority`:
  ```javascript
  db.balancer_events.createIndex({ "status": 1, "priority": -1 });
  ```
  Optimizes fetching of pending events sorted by priority.
- **TTL index** on `createdAt` (events older than 7 days will be removed automatically):
  ```javascript
  db.balancer_events.createIndex({ "createdAt": 1 }, { expireAfterSeconds: 604800 });
  ```

### 2. `balancing_algorithms`
Stores details of load balancing algorithms.

#### Schema
```json
{
  "_id": "ObjectId",
  "name": "string",
  "description": "string",
  "createdAt": "timestamp",
  "isActive": "boolean" // Whether this algorithm is currently in use
}
```

#### Indexes
- **Single index** on `isActive`:
  ```javascript
  db.balancing_algorithms.createIndex({ "isActive": 1 });
  ```
  Facilitates quick lookup of the active algorithm.

### 3. `balancer_results`
Stores results of executed balancing operations.

#### Schema
```json
{
  "_id": "ObjectId",
  "algorithmId": "ObjectId",
  "eventId": "ObjectId",
  "result": "object", // Details of the balancing result
  "executedAt": "timestamp"
}
```

#### Indexes
- **Single index** on `algorithmId`:
  ```javascript
  db.balancer_results.createIndex({ "algorithmId": 1 });
  ```
  Speeds up queries for results by algorithm.
- **Compound index** on `eventId` and `executedAt`:
  ```javascript
  db.balancer_results.createIndex({ "eventId": 1, "executedAt": -1 });
  ```
  Optimizes retrieving results for specific events sorted by execution time.

## Key Selection and Query Patterns

### Key Selections:
- **Primary Key (_id)**: Automatically indexed and used for unique identification of documents.
- **Indexes**: 
  - `balancer_events`: Compound on `status` and `priority`, TTL on `createdAt`.
  - `balancing_algorithms`: Single on `isActive`.
  - `balancer_results`: Compound on `eventId` and `executedAt`, single on `algorithmId`.

### Query Patterns:
1. **Retrieve pending events sorted by priority**:
   ```javascript
   db.balancer_events.find({ "status": "pending" }).sort({ "priority": -1 });
   ```

2. **Get active balancing algorithm**:
   ```javascript
   db.balancing_algorithms.findOne({ "isActive": true });
   ```

3. **Fetch results of a specific event**:
   ```javascript
   db.balancer_results.find({ "eventId": "<EVENT_ID>" }).sort({ "executedAt": -1 });
   ```

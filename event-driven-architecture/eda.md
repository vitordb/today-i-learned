# Events

* Something that happened in the past
* Usually it will leave collateral effects.
* Can work internally in the software or externally.
* Domain events: Something that happened in the application, specifically in the domain logic, where you have your business logic. Ex: aggregates.

## Types of events

1 - Event Notification: Shortest type of notification. Ex: {"order": 1, "status": "approved"}

2 - Event Carried State Transfer: Complete type of notification, data stream. Ex: {"order": 1, "products": [{}], "price": 10.00, "tax": 0.01 }

3 - Event sourcing: 

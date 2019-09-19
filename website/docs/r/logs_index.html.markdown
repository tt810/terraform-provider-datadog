---
layout: "datadog"
page_title: "Datadog: datadog_logs_index"
sidebar_current: "docs-datadog-resource-logs-index"
description: |-
  Provides a Datadog logs index resource. This can be used to create and manage logs indexes.
---

# datadog_logs_index

Provides a Datadog [Logs Index API](https://docs.datadoghq.com/api/?lang=python#logs-indexes) resource. This can be used to create and manage Datadog logs indexes.

## Example Usage: Datadog logs index

```hcl
# Update a Datadog logs index
resource "datadog_logs_index" "sample_index" {
    name = "your index"
    
}
```
## Example Usage: Datadog logs pipeline order

```hcl
resource "datadog_logs_indexorder" "sample_index_order" {
    name = "sample_index_order"
    depends_on = [
        "datadog_logs_index.sample_index"
    ]
    pipelines = [
        "${datadog_logs_index.sample_index.id}"
    ]
}
```


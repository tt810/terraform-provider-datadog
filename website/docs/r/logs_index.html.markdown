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
    filter {
        query = "*"
    }
    exclusion_filter {
        name = "Filter coredns logs"
        is_enabled = true
        filter {
            query = "app:coredns"
            sample_rate = 0.97
        }
    }
    exclusion_filter {
        name = "Kubernetes apiserver"
        is_enabled = true
        filter {
            query = "service:kube_apiserver"
            sample_rate = 1.0
        }
    }
}
```
## Example Usage: Datadog logs pipeline order

```hcl
resource "datadog_logs_indexorder" "sample_index_order" {
    name = "sample_index_order"
    depends_on = [
        "datadog_logs_index.sample_index"
    ]
    indexes = [
        "${datadog_logs_index.sample_index.id}"
    ]
}
```

## Argument Reference

The following arguments are supported in resource `datadog_logs_index`:
* `name` - (Required) The name of the index.
* `filter` - (Required) Logs filter.
  * `query` - (Required) Logs filter criteria. Only logs matching this filter criteria are considered for this index.
* `exclusion_filter` - (Required) List of exclusion filter.
  * `name` - (Optional) The name of the exclusion filter.
  * `is_enabled` - (Optional, default = false) A boolean stating if the exclusion is active or not.
  * `filter` - (Required)
    * `query` - (Required) Only logs matching the filter criteria and the query of the parent index will be considered for this exclusion filter.
    * `sample_rate` - (Optional, default = 0.0) the fraction of logs excluded by the exclusion filter, when active.

The following arguments are supported in resource `datadog_logs_indexorder`:
* `name` - (Required) The unique name of the index order. 
* `indexes` - (Required) The index resource list. Logs are tested against the query filter of each index one by one, following the order of the list.

## Import

For the existing indexes, do `terraform import <resource.name> <indexName>` to import them to terraform.  

## Important Notes

Each `datadog_logs_order` resource defines a complete index. The order of index are maintained in resource `datadog_logs_indexorder`.
There should be just one `datadog_logs_indexorder` resource. The creation and deletion of logs index and index order are currently
not supported through terraform provider.

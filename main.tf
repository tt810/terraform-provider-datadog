resource "datadog_logs_pipeline" "test_import" {
	name = "imported pipeline"
	is_enabled = false
	filter {
		query = "source:kafka"
	}
}

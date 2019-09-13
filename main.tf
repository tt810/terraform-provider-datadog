resource "datadog_logs_pipeline" "my" {
	name = "my pipeline"
	is_enabled = true
	filter {
		query = "source:kafka"
	}
	processor {
		arithmetic_processor {
			name = "updated arithmetic processor"
			is_enabled = true
			expression = "(time1 - time2)*1000"
			target = "my_arithmetic"
			is_replace_missing = true
		}
	}
	processor {
        attribute_remapper {
            name = "test attribute processor"
            is_enabled = true
            sources = ["db.instance"]
            source_type = "tag"
            target = "db"
            target_type = "tag"
        }
    }

}


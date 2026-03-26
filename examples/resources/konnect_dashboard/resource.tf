resource "konnect_dashboard" "my_dashboard" {
  provider = konnect-beta
  definition = {
    preset_filters = [
      {
        field    = "ai_provider"
        operator = "not_in"
        value    = "{ \"see\": \"documentation\" }"
      }
    ]
    tiles = [
      {
        chart = {
          definition = {
            chart = {
              choropleth_map = {
                chart_title = "...my_chart_title..."
                type        = "choropleth_map"
              }
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = [
                  "status_code_grouped",
                ]
                filters = [
                  {
                    field    = "realm"
                    operator = "not_empty"
                    value    = "{ \"see\": \"documentation\" }"
                  }
                ]
                granularity = "twelveHourly"
                metrics = [
                  "kong_latency_average"
                ]
                time_range = {
                  relative = {
                    time_range = "1h"
                    type       = "relative"
                    tz         = "Etc/UTC"
                  }
                }
              }
            }
          }
          layout = {
            position = {
              col = 4
              row = 5
            }
            size = {
              cols = 6
              rows = 8
            }
          }
          type = "chart"
        }
      }
    ]
  }
  labels = {
    key = "value"
  }
  name = "...my_name..."
}
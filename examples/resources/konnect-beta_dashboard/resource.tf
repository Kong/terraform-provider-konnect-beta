resource "konnect-beta_dashboard" "my_dashboard" {
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
              donut = {
                chart_title = "...my_chart_title..."
                type        = "donut"
              }
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = [
                  "api_product"
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
                  "kong_latency_p50"
                ]
                time_range = {
                  relative = {
                    time_range = "current_week"
                    type       = "relative"
                    tz         = "...my_tz..."
                  }
                }
              }
              llm_usage = {
                datasource = "llm_usage"
                dimensions = [
                  "consumer"
                ]
                filters = [
                  {
                    field    = "application"
                    operator = "empty"
                    value    = "{ \"see\": \"documentation\" }"
                  }
                ]
                granularity = "tenMinutely"
                metrics = [
                  "ai_request_count"
                ]
                time_range = {
                  absolute = {
                    end   = "2022-11-26T07:30:44.592Z"
                    start = "2022-01-09T02:25:36.303Z"
                    type  = "absolute"
                    tz    = "...my_tz..."
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
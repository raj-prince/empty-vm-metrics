from google.cloud import monitoring_v3
import sys

def sample_list_time_series(start_time_sec, end_time_sec, period) -> None:
    # Create a client
    client = monitoring_v3.MetricServiceClient()


    metric_type = "custom.googleapis.com/opencensus/test_star_count/star_count"
    #metric_type = "custom.googleapis.com/gcsfuse/gcs/read_bytes_count"

    metric_filter = ('metric.type = "{metric_type}"').format(metric_type=metric_type)


    #period = (end_time_sec - start_time_sec)


    print("Start time: ", start_time_sec)
    print("End time: ", end_time_sec)
    print("Period: ", period)
    interval = monitoring_v3.TimeInterval(
        end_time={'seconds': int(end_time_sec)},
        start_time={'seconds': int(start_time_sec)})

    aggregation_peak = monitoring_v3.Aggregation(
        alignment_period={'seconds': period},
        per_series_aligner=monitoring_v3.Aggregation.Aligner
        .ALIGN_DELTA,
    )


    try:
    # Initialize request argument(s)
        request = monitoring_v3.ListTimeSeriesRequest(
            #name="projects/gcs-fuse-test-ml/custom.googleapis.com/gcsfuse",
            name="projects/gcs-fuse-test-ml/test_star_count",
            filter=metric_filter,
            view="FULL",
            interval=interval,
            aggregation=aggregation_peak,
        )
    except:
        raise GoogleAPICallError("Testing")


    print(request)
    # Make the request
    page_result = client.list_time_series(request=request)

    if len(page_result) == 0:
        print("Empty result")
        exit(1)

    # Handle the response
    for response in page_result:
      if response.value == 0:
        print("Empty result")
        exit(1)
        print(response)

def main() -> None:
  if len(sys.argv) != 4:
    raise Exception('Invalid arguments.')
  start_time_sec = int(sys.argv[1])
  end_time_sec = int(sys.argv[2])
  period = int(sys.argv[3])
  sample_list_time_series(start_time_sec, end_time_sec, period)

if __name__ == '__main__':
  main()


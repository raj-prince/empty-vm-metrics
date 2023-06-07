// Example of exporting a custom metric from OpenCensus to Cloud Monitoring.
package main
 
import (
        "context"
        "fmt"
        "log"
        "time"
 
        "contrib.go.opencensus.io/exporter/stackdriver"
        "go.opencensus.io/stats"
        "go.opencensus.io/stats/view"
//        "golang.org/x/exp/rand"
)
 
var (
        // The restaurant rating in number of stars.
        starCount = stats.Int64("star_count", "The star rating for the restaurant (0-5)", "stars")
)
 
func main() {
        ctx := context.Background()
 
        // Register the view. It is imperative that this step exists,
        // otherwise recorded metrics will be dropped and never exported.
        v := &view.View{
                Name:        "star_count",
                Measure:     starCount,
                Description: "Restaurant star rating 0-5",
                Aggregation: view.Sum(),
        }
        if err := view.Register(v); err != nil {
                log.Fatalf("Failed to register the view: %v", err)
        }
 
        // Enable OpenCensus exporters to export metrics
        // to Google Cloud Monitoring.
        // Exporters use Application Default Credentials to authenticate.
        // See https://developers.google.com/identity/protocols/application-default-credentials
        // for more details.
        // The Stackdriver Exporter sets a default MonitoredResource of type “global”
        exporter, err := stackdriver.NewExporter(stackdriver.Options{
            // ProjectID <change this value>
            ProjectID: "gcs-fuse-test-ml",
            // MetricPrefix helps uniquely identify your metrics. <change this value>
            MetricPrefix: "test_star_count", 
            // ReportingInterval sets the frequency of reporting metrics
            // to the Cloud Monitoring backend.
            ReportingInterval: 30 * time.Second,
        })
 
        if err != nil {
                log.Fatal(err)
        }
        // Flush must be called before main() exits to ensure metrics are recorded.
        defer exporter.Flush()
 
        if err := exporter.StartMetricsExporter(); err != nil {
                log.Fatalf("Error starting metric exporter: %v", err)
        }
        defer exporter.StopMetricsExporter()
 
        // Record 100 fake count values between 0-5.
      //  rand.Intn(n) returns a random number between 0 and n - 1
        for i := 0; i < 14; i++ {

                time.Sleep(5 * time.Second)
        //        random_star_count := int64(rand.Intn(6))
		random_star_count := int64(1)
                fmt.Println("Star count: ", random_star_count)
                stats.Record(ctx, starCount.M(random_star_count))
                
		fmt.Println("Wait for 1 minute")
                // Wait 1 second until we write the next count
        }
 
        fmt.Println("Done recording metrics")
}

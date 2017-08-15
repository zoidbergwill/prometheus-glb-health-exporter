package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

const containerEngineInstanceGroup = "https://www.googleapis.com/compute/v1/projects/%s/zones/%s/instanceGroups/k8s-ig"

func main() {
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if project == "" {
		log.Fatal("You need to set env var: GOOGLE_CLOUD_PROJECT")
	}

	region := os.Getenv("GOOGLE_CLOUD_ZONE")
	if region == "" {
		log.Fatal("You need to set env var: GOOGLE_CLOUD_PROJECT")
	}

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := computeService.BackendServices.List(project)
	if err := req.Pages(ctx, func(page *compute.BackendServiceList) error {
		for _, backendService := range page.Items {

			rb := &compute.ResourceGroupReference{
				Group: fmt.Sprintf(containerEngineInstanceGroup, project, region),
			}

			resp, err := computeService.BackendServices.GetHealth(project, backendService.Name, rb).Context(ctx).Do()
			if err != nil {
				log.Fatal(err)
			}
			total := 0
			healthy := 0
			for _, individualHealthStatus := range resp.HealthStatus {
				total += 1
				if individualHealthStatus.HealthState == "HEALTHY" {
					healthy += 1
				}
			}
			fmt.Printf("%s: %v/%v\n", backendService.Name, healthy, total)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

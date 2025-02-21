package releases

import (
	"context"
	"fmt"
	"slices"
)

func ExampleClient_Products() {
	client, err := New()
	if err != nil {
		panic(err)
	}

	products, err := client.Products(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Contains Terraform?", slices.Contains(products, "terraform"))
	fmt.Println("Contains Packer?", slices.Contains(products, "packer"))
	// Output:
	// Contains Terraform? true
	// Contains Packer? true
}

func ExampleClient_ReleasesPaged() {
	client, err := New()
	if err != nil {
		panic(err)
	}

	// Note that Waypoint as a released binary is discontinued, so the output should be stable even
	// though we are hitting the production endpoint.
	releasePages, err := client.ReleasesPaged(context.Background(), "waypoint", LicenseClassOSS)
	if err != nil {
		panic(err)
	}

	pageNum := 0
	for page, err := range releasePages {
		if err != nil {
			panic(err)
		}
		pageNum++

		fmt.Printf("Page %d has %d items\n", pageNum, len(page))
	}

	// Output:
	// Page 1 has 16 items
	// Page 2 has 16 items
	// Page 3 has 11 items
}

func ExampleClient_Releases() {
	client, err := New()
	if err != nil {
		panic(err)
	}

	// Note that Waypoint as a released binary is discontinued, so the output should be stable even
	// though we are hitting the production endpoint.
	releases, err := client.Releases(context.Background(), "waypoint", LicenseClassOSS)
	if err != nil {
		panic(err)
	}

	for release, err := range releases {
		if err != nil {
			panic(err)
		}

		fmt.Println("Release", release.Version)
	}

	// Output:
	// Release 0.11.4
	// Release 0.11.3
	// Release 0.11.2
	// Release 0.11.1
	// Release 0.11.0
	// Release 0.10.5
	// Release 0.10.4
	// Release 0.10.3
	// Release 0.10.2
	// Release 0.10.1
	// Release 0.10.0
	// Release 0.9.1
	// Release 0.9.0
	// Release 0.8.2
	// Release 0.8.1
	// Release 0.8.0
	// Release 0.7.2
	// Release 0.7.1
	// Release 0.7.0
	// Release 0.6.3
	// Release 0.6.2
	// Release 0.6.1
	// Release 0.6.0
	// Release 0.5.2
	// Release 0.5.1
	// Release 0.5.0
	// Release 0.4.2
	// Release 0.4.1
	// Release 0.4.0
	// Release 0.3.2
	// Release 0.3.1
	// Release 0.3.0
	// Release 0.2.4
	// Release 0.2.3
	// Release 0.2.2
	// Release 0.2.1
	// Release 0.2.0
	// Release 0.1.5
	// Release 0.1.4
	// Release 0.1.3
	// Release 0.1.2
	// Release 0.1.1
	// Release 0.1.0
}

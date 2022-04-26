// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccApigeeInstance_apigeeInstanceBasicTestExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
		"random_suffix":   randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApigeeInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeInstance_apigeeInstanceBasicTestExample(context),
			},
			{
				ResourceName:            "google_apigee_instance.apigee_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ip_range", "org_id"},
			},
		},
	})
}

func testAccApigeeInstance_apigeeInstanceBasicTestExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_instance" "apigee_instance" {
  name     = "tf-test%{random_suffix}"
  location = "us-central1"
  org_id   = google_apigee_organization.apigee_org.id
}
`, context)
}

func TestAccApigeeInstance_apigeeInstanceCidrRangeTestExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
		"random_suffix":   randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApigeeInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeInstance_apigeeInstanceCidrRangeTestExample(context),
			},
			{
				ResourceName:            "google_apigee_instance.apigee_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ip_range", "org_id"},
			},
		},
	})
}

func testAccApigeeInstance_apigeeInstanceCidrRangeTestExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 22
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_instance" "apigee_instance" {
  name     = "tf-test%{random_suffix}"
  location = "us-central1"
  org_id   = google_apigee_organization.apigee_org.id
  peering_cidr_range = "SLASH_22"
}
`, context)
}

func TestAccApigeeInstance_apigeeInstanceIpRangeTestExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
		"random_suffix":   randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApigeeInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeInstance_apigeeInstanceIpRangeTestExample(context),
			},
			{
				ResourceName:            "google_apigee_instance.apigee_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ip_range", "org_id"},
			},
		},
	})
}

func testAccApigeeInstance_apigeeInstanceIpRangeTestExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 22
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_instance" "apigee_instance" {
  name     = "tf-test%{random_suffix}"
  location = "us-central1"
  org_id   = google_apigee_organization.apigee_org.id
  ip_range = "10.87.8.0/22"
}
`, context)
}

func TestAccApigeeInstance_apigeeInstanceServiceAttachmentBasicTestExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
		"random_suffix":   randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApigeeInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeInstance_apigeeInstanceServiceAttachmentBasicTestExample(context),
			},
			{
				ResourceName:            "google_apigee_instance.apigee_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ip_range", "org_id"},
			},
		},
	})
}

func testAccApigeeInstance_apigeeInstanceServiceAttachmentBasicTestExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_compute_address" "psc_ilb_consumer_address" {
  name   = "psc-ilb-consumer-address"
  region = "us-west2"

  subnetwork   = "default"
  address_type = "INTERNAL"

  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_forwarding_rule" "psc_ilb_consumer" {
  name   = "psc-ilb-consumer-forwarding-rule"
  region = "us-west2"

  target                = google_compute_service_attachment.psc_ilb_service_attachment.id
  load_balancing_scheme = "" # need to override EXTERNAL default when target is a service attachment
  network               = "default"
  ip_address            = google_compute_address.psc_ilb_consumer_address.id

  project = google_project.project.project_id
}

resource "google_compute_forwarding_rule" "psc_ilb_target_service" {
  name   = "producer-forwarding-rule"
  region = "us-west2"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.producer_service_backend.id
  all_ports             = true
  network               = google_compute_network.psc_ilb_network.name
  subnetwork            = google_compute_subnetwork.psc_ilb_producer_subnetwork.name

  project = google_project.project.project_id
}

resource "google_compute_region_backend_service" "producer_service_backend" {
  name   = "producer-service"
  region = "us-west2"

  health_checks = [google_compute_health_check.producer_service_health_check.id]

  project = google_project.project.project_id
}

resource "google_compute_health_check" "producer_service_health_check" {
  name = "producer-service-health-check"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }

  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_network" "psc_ilb_network" {
  name = "psc-ilb-network"
  auto_create_subnetworks = false

  project     = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_subnetwork" "psc_ilb_producer_subnetwork" {
  name   = "psc-ilb-producer-subnetwork"
  region = "us-west2"

  network       = google_compute_network.psc_ilb_network.id
  ip_cidr_range = "10.0.0.0/16"

  project = google_project.project.project_id
}

resource "google_compute_subnetwork" "psc_ilb_nat" {
  name   = "psc-ilb-nat"
  region = "us-west2"

  network       = google_compute_network.psc_ilb_network.id
  purpose       =  "PRIVATE_SERVICE_CONNECT"
  ip_cidr_range = "10.1.0.0/16"

  project = google_project.project.project_id
}

resource "google_compute_service_attachment" "psc_ilb_service_attachment" {
  name        = "my-psc-ilb"
  region      = "us-west2"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol    = true
  connection_preference    = "ACCEPT_AUTOMATIC"
  nat_subnets              = [google_compute_subnetwork.psc_ilb_nat.id]
  target_service           = google_compute_forwarding_rule.psc_ilb_target_service.id

  project = google_project.project.project_id
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_instance" "apigee_instance" {
  name                 = "tf-test%{random_suffix}"
  location             = "us-central1"
  org_id               = google_apigee_organization.apigee_org.id
  consumer_accept_list = [123456, google_project.project.number]
}
`, context)
}

func testAccCheckApigeeInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_apigee_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ApigeeBasePath}}{{org_id}}/instances/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ApigeeInstance still exists at %s", url)
			}
		}

		return nil
	}
}

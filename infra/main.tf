# Configure the Google Cloud provider
provider "google" {
	project = "your-gcp-project-id"  # Replace with your GCP project ID
	region  = "us-west1"             # Set the default region for resources
}

# Define a Google Compute Engine instance resource
resource "google_compute_instance" "example" {
	name         = "example-instance"  # Name of the instance
	machine_type = "f1-micro"          # Machine type (instance size)
	zone         = "us-west1-a"        # Zone where the instance will be created

	# Configure the boot disk for the instance
	boot_disk {
		initialize_params {
			image = "debian-cloud/debian-9"  # Image to use for the boot disk
		}
	}

	# Configure the network interface for the instance
	network_interface {
		network = "default"  # Use the default network
		access_config {      # Enable external access (assign an external IP)
		}
	}

	# Add tags to the instance
	tags = ["example-instance"]  # Tags for the instance
}
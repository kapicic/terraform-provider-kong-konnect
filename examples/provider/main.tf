terraform {
  required_providers {
    liblab = {
      source = "hashicorp/liblab"
    }
  }
}

provider "liblab" {

  host = "http://localhost:8001/"

  verify_tls = true

}

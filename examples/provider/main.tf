terraform {
  required_providers {
    kong = {
      source = "hashicorp/kong"
    }
  }
}

provider "kong" {

  host = "http://localhost:8001/"

  verify_tls = true

}

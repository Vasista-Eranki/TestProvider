terraform {
  required_providers {
    helloworld = {
      source = "app.terraform.io/Hexagon-PPM/helloworld"
    }
  }
}

provider "helloworld" {
  //plugin_dir = "c:\\users\\vseranki\\gobin"
}
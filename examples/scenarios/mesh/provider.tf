terraform {
  required_providers {
    konnect = {
      source = "kong/konnect"
    }
    konnect-beta = {
      source = "kong/konnect-beta"
    }
    time = {
      source  = "hashicorp/time"
      version = "0.12.1"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.35.1"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "2.17.0"
    }
  }
}

variable "konnect_personal_access_token" {
  type    = string
}

variable "region" {
  type    = string
  default = "us"
}

variable "k8s_cluster_config_path" {
  type        = string
  description = "The location where this cluster's kubeconfig will be saved to."
  default     = ""
}

locals {
  should_run = var.k8s_cluster_config_path != ""
}


provider "time" {}

provider "konnect" {
  personal_access_token = var.konnect_personal_access_token
  server_url            = "https://${var.region}.api.konghq.com"
}


provider "konnect-beta" {
  personal_access_token = var.konnect_personal_access_token
  server_url            = "https://${var.region}.api.konghq.com"
}

provider "kubernetes" {
  config_path = pathexpand(var.k8s_cluster_config_path)
}

provider "helm" {
  kubernetes {
    config_path = pathexpand(var.k8s_cluster_config_path)
  }
}

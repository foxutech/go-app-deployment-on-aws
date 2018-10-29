variable "aws_access_key" {
  description = "This name will prefix all resources, and be added as the value for the 'Name' tag where supported"
  type = "string"
}

variable "aws_secret_key" {
  description = "This name will prefix all resources, and be added as the value for the 'Name' tag where supported"
  type = "string"
}

variable "aws_region" {
  description = "AWS region on which we will setup the swarm cluster"
  default = "us-west-2"
}
 
variable "instance_type" {
  description = "Instance type"
  default = "t2.micro"
}
 
variable "key_path" {
  description = "SSH Public Key path"
  default = "./test.pem.pub"
}
 
variable "bootstrap_path" {
  description = "Script to install Docker Engine"
  default = "./docker.sh"
}

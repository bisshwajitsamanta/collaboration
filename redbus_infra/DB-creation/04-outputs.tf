output "db_instance_password" {
  value = nonsensitive(module.db.db_instance_password)

}

output "db_instance_id" {
  description = "The RDS instance ID"
  value       = module.db.db_instance_id
}

output "db_instance_name" {
  description = "The database name"
  value       = module.db.db_instance_name
}

output "db_instance_username" {
  description = "The master username for the database"
  value       = nonsensitive(module.db.db_instance_username)
}

output "db_instance_port" {
  description = "The database port"
  value       = module.db.db_instance_port
}

output "db_instance_endpoint" {
  value = module.db.db_instance_endpoint
}

output "db_instance_resource_id" {
  value = module.db.db_instance_resource_id
}

output "database_internet_gateway_route_id" {
  value = module.vpc.database_internet_gateway_route_id
}

output "default_route_table_id" {
  value = module.vpc.default_route_table_id
}
output "igw_id" {
  value = module.vpc.igw_id
}
output "public_internet_gateway_route_id" {
  value = module.vpc.public_internet_gateway_route_id
}
output "public_route_table_ids" {
  value = module.vpc.public_route_table_ids
}
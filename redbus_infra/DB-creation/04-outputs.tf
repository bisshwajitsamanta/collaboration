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
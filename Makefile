dry-migrate: ## Try migration
	mysqldef -u user -p password -h 127.0.0.1 -P 3306 db --dry-run < ./_tools/mysql/schema.sql

migrate:  ## Execute migration
	mysqldef -u user -p password -h 127.0.0.1 -P 3306 db < ./_tools/mysql/schema.sql
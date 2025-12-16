-- Drop indexes (MySQL syntax)
DROP INDEX IF EXISTS idx_service_charge_statuses_service_charge_id ON service_charge_statuses;
DROP INDEX IF EXISTS idx_service_charge_items_service_charge_id ON service_charge_items;
DROP INDEX IF EXISTS idx_service_charges_account_id ON service_charges;
DROP INDEX IF EXISTS idx_one_shot_usages_catalog_id ON one_shot_usages;
DROP INDEX IF EXISTS idx_one_shot_usages_account_id ON one_shot_usages;
DROP INDEX IF EXISTS idx_subscriptions_next_billing_date ON subscriptions;
DROP INDEX IF EXISTS idx_subscriptions_catalog_id ON subscriptions;
DROP INDEX IF EXISTS idx_subscriptions_account_id ON subscriptions;
DROP INDEX IF EXISTS idx_catalog_prices_catalog_id ON catalog_prices;
DROP INDEX IF EXISTS idx_catalogs_tenant_id ON catalogs;
DROP INDEX IF EXISTS idx_accounts_tenant_id ON accounts;

-- Drop tables in reverse order (respecting foreign key constraints)
DROP TABLE IF EXISTS service_charge_statuses;
DROP TABLE IF EXISTS service_charge_items;
DROP TABLE IF EXISTS service_charges;
DROP TABLE IF EXISTS one_shot_usages;
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS catalog_prices;
DROP TABLE IF EXISTS catalogs;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS tenants;


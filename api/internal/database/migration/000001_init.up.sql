-- Create tenants table
CREATE TABLE IF NOT EXISTS tenants (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL PRIMARY KEY,
    external_account_id VARCHAR(255) NOT NULL,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_accounts_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    CONSTRAINT uq_accounts_external_tenant UNIQUE (external_account_id, tenant_id)
);

-- Create catalogs table
CREATE TABLE IF NOT EXISTS catalogs (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_catalogs_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);

-- Create catalog_prices table
CREATE TYPE billing_type AS ENUM ('ONE_SHOT', 'RECURRING');
CREATE TYPE billing_cycle AS ENUM ('MONTHLY', 'YEARLY', 'QUARTERLY');

CREATE TABLE IF NOT EXISTS catalog_prices (
    id BIGSERIAL PRIMARY KEY,
    catalog_id BIGINT NOT NULL,
    amount BIGINT NOT NULL,
    billing_type billing_type NOT NULL,
    billing_cycle billing_cycle,
    start_date DATE NOT NULL,
    end_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_catalog_prices_catalog FOREIGN KEY (catalog_id) REFERENCES catalogs(id)
);

-- Create subscriptions table
CREATE TABLE IF NOT EXISTS subscriptions (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    catalog_id BIGINT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    idempotency_key VARCHAR(255) NOT NULL UNIQUE,
    next_billing_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_subscriptions_account FOREIGN KEY (account_id) REFERENCES accounts(id),
    CONSTRAINT fk_subscriptions_catalog FOREIGN KEY (catalog_id) REFERENCES catalogs(id)
);

-- Create one_shot_usages table
CREATE TABLE IF NOT EXISTS one_shot_usages (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    catalog_id BIGINT NOT NULL,
    idempotency_key VARCHAR(255) NOT NULL UNIQUE,
    used_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_one_shot_usages_account FOREIGN KEY (account_id) REFERENCES accounts(id),
    CONSTRAINT fk_one_shot_usages_catalog FOREIGN KEY (catalog_id) REFERENCES catalogs(id)
);

-- Create service_charges table
CREATE TABLE IF NOT EXISTS service_charges (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    amount BIGINT NOT NULL,
    latest_status_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_service_charges_account FOREIGN KEY (account_id) REFERENCES accounts(id)
);

-- Create service_charge_items table
CREATE TABLE IF NOT EXISTS service_charge_items (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    service_charge_id BIGINT NOT NULL,
    one_shot_usage_id BIGINT,
    subscription_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_service_charge_items_charge FOREIGN KEY (service_charge_id) REFERENCES service_charges(id),
    CONSTRAINT fk_service_charge_items_one_shot FOREIGN KEY (one_shot_usage_id) REFERENCES one_shot_usages(id),
    CONSTRAINT fk_service_charge_items_subscription FOREIGN KEY (subscription_id) REFERENCES subscriptions(id)
);

-- Create service_charge_statuses table
CREATE TYPE service_charge_status AS ENUM ('RESERVED', 'COMPLETED', 'FAILED', 'CANCELLED');

CREATE TABLE IF NOT EXISTS service_charge_statuses (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status service_charge_status NOT NULL,
    service_charge_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_service_charge_statuses_charge FOREIGN KEY (service_charge_id) REFERENCES service_charges(id)
);

-- Add foreign key constraint for latest_status_id after service_charge_statuses table is created
ALTER TABLE service_charges
    ADD CONSTRAINT fk_service_charges_latest_status
    FOREIGN KEY (latest_status_id) REFERENCES service_charge_statuses(id);

-- Create indexes for better query performance
CREATE INDEX idx_accounts_tenant_id ON accounts(tenant_id);
CREATE INDEX idx_catalogs_tenant_id ON catalogs(tenant_id);
CREATE INDEX idx_catalog_prices_catalog_id ON catalog_prices(catalog_id);
CREATE INDEX idx_subscriptions_account_id ON subscriptions(account_id);
CREATE INDEX idx_subscriptions_catalog_id ON subscriptions(catalog_id);
CREATE INDEX idx_subscriptions_next_billing_date ON subscriptions(next_billing_date);
CREATE INDEX idx_one_shot_usages_account_id ON one_shot_usages(account_id);
CREATE INDEX idx_one_shot_usages_catalog_id ON one_shot_usages(catalog_id);
CREATE INDEX idx_service_charges_account_id ON service_charges(account_id);
CREATE INDEX idx_service_charge_items_service_charge_id ON service_charge_items(service_charge_id);
CREATE INDEX idx_service_charge_statuses_service_charge_id ON service_charge_statuses(service_charge_id);


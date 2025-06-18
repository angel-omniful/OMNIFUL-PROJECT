CREATE TABLE IF NOT EXISTS inventory (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id TEXT NOT NULL,
    seller_id TEXT NOT NULL,
    hub_id UUID NOT NULL,
    sku_id UUID NOT NULL,
    quantity INT DEFAULT 0,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (tenant_id, seller_id, hub_id, sku_id)
);
CREATE INDEX IF NOT EXISTS idx_inventory_tenant_id ON inventory (tenant_id);
CREATE INDEX IF NOT EXISTS idx_inventory_seller_id ON inventory (seller_id);
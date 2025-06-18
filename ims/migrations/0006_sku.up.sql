CREATE TABLE IF NOT EXISTS skus (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id TEXT NOT NULL,
    seller_id TEXT NOT NULL,
    sku_code TEXT NOT NULL,
    name TEXT,
    category TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (tenant_id, seller_id, sku_code)
);
CREATE INDEX IF NOT EXISTS idx_skus_tenant_id ON skus (tenant_id);
CREATE INDEX IF NOT EXISTS idx_skus_seller_id ON skus (seller_id);  
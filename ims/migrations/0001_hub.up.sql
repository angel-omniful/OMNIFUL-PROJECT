CREATE TABLE IF NOT EXISTS hubs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id TEXT NOT NULL,
    seller_id TEXT NOT NULL,
    name TEXT NOT NULL,
    location TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_hubs_tenant_id ON hubs (tenant_id);
CREATE INDEX IF NOT EXISTS idx_hubs_seller_id ON hubs (seller_id);
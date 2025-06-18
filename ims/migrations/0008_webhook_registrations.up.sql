CREATE TABLE IF NOT EXISTS webhook_registrations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id TEXT NOT NULL,
    url TEXT NOT NULL,
    event_type TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_webhook_registrations_tenant_id ON webhook_registrations (tenant_id);
CREATE INDEX IF NOT EXISTS idx_webhook_registrations_event_type ON webhook_registrations (event_type);
CREATE INDEX IF NOT EXISTS idx_webhook_registrations_created_at ON webhook_registrations (created_at);
CREATE INDEX IF NOT EXISTS idx_webhook_registrations_is_active ON webhook_registrations (is_active);     
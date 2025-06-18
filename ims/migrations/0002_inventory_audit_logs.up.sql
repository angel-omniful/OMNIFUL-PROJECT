CREATE TABLE IF NOT EXISTS inventory_audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    inventory_id UUID NOT NULL,
    change INT NOT NULL,
    reason TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_inventory_audit_logs_inventory_id ON inventory_audit_logs (inventory_id);
CREATE INDEX IF NOT EXISTS idx_inventory_audit_logs_created_at ON inventory_audit_logs (created_at);    
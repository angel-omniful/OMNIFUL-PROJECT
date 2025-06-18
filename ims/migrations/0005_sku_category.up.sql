CREATE TABLE IF NOT EXISTS sku_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE
);
CREATE INDEX IF NOT EXISTS idx_sku_categories_name ON sku_categories (name);
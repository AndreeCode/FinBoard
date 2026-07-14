CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID
);

CREATE TABLE IF NOT EXISTS permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID
);

CREATE TABLE IF NOT EXISTS role_permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id UUID REFERENCES roles(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID,
    last_login TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    parent_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    amount DECIMAL(12,2) NOT NULL,
    type VARCHAR(20) NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    received_date TIMESTAMP,
    due_date TIMESTAMP,
    canceled BOOLEAN DEFAULT FALSE,
    description TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID
);

CREATE TABLE IF NOT EXISTS credits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    person_name VARCHAR(100) NOT NULL,
    amount DECIMAL(12,2) NOT NULL,
    interest_rate DECIMAL(5,2) DEFAULT 0,
    is_creditor BOOLEAN NOT NULL DEFAULT TRUE,
    is_secure BOOLEAN NOT NULL DEFAULT TRUE,
    due_date TIMESTAMP,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID
);

CREATE TABLE IF NOT EXISTS investments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID REFERENCES transactions(id) ON DELETE CASCADE,
    expected_gain DECIMAL(12,2),
    risk_level VARCHAR(20),
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID
);

CREATE TABLE IF NOT EXISTS audit (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    table_name VARCHAR(100) NOT NULL,
    record_id UUID,
    action VARCHAR(10) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    changed_by UUID REFERENCES users(id),
    changed_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION audit_trigger_fn() RETURNS TRIGGER AS $$
DECLARE
    user_id UUID;
BEGIN
    user_id := current_setting('app.current_user', true)::UUID;
    IF user_id IS NULL AND TG_TABLE_NAME = 'users' AND TG_OP = 'INSERT' THEN
        user_id := NEW.id;
    END IF;

    IF TG_OP = 'INSERT' THEN
        INSERT INTO audit(table_name, record_id, action, new_data, changed_by)
        VALUES (TG_TABLE_NAME, NEW.id, TG_OP, to_jsonb(NEW), user_id);
        RETURN NEW;
    ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO audit(table_name, record_id, action, old_data, new_data, changed_by)
        VALUES (TG_TABLE_NAME, NEW.id, TG_OP, to_jsonb(OLD), to_jsonb(NEW), user_id);
        RETURN NEW;
    ELSIF TG_OP = 'DELETE' THEN
        INSERT INTO audit(table_name, record_id, action, old_data, changed_by)
        VALUES (TG_TABLE_NAME, OLD.id, TG_OP, to_jsonb(OLD), user_id);
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;

DO $$
DECLARE
    tbl TEXT;
BEGIN
    FOR tbl IN SELECT table_name FROM information_schema.tables
               WHERE table_schema='public' AND table_name IN ('users','roles','permissions','role_permissions','categories','transactions','credits','investments')
    LOOP
        IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = tbl || '_audit') THEN
            EXECUTE format('CREATE TRIGGER %I_audit AFTER INSERT OR UPDATE OR DELETE ON %I FOR EACH ROW EXECUTE FUNCTION audit_trigger_fn();', tbl, tbl);
        END IF;
    END LOOP;
END$$;

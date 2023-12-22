create table role_access (
    id uuid primary key,
    role_id uuid not null,
    module_id uuid,
    tenant_id uuid,
    permission jsonb,
    status status_enum default 'active'::status_enum,
    created_at timestamptz default current_timestamp,
    created_by uuid,
    updated_at timestamptz,
    updated_by uuid,
    deleted_at timestamptz,
    deleted_by uuid,
    doc_version int default 0,
    unique(tenant_id, role_id, module_id),
    foreign key (tenant_id)
        references tenants(id)
        on delete set null
        on update cascade,
    foreign key (role_id)
        references roles(id)
        on delete set null
        on update cascade,
    foreign key (module_id)
        references modules(id)
        on delete set null
        on update cascade
);

CREATE TRIGGER role_access_update_trigger
BEFORE UPDATE ON role_access
FOR EACH ROW
WHEN (
    OLD.role_id IS DISTINCT FROM NEW.role_id OR
    OLD.module_id IS DISTINCT FROM NEW.module_id OR
    OLD.tenant_id IS DISTINCT FROM NEW.tenant_id OR
    OLD.permission IS DISTINCT FROM NEW.permission OR
    OLD.status IS DISTINCT FROM NEW.status
)
EXECUTE FUNCTION update_doc_and_last_update();

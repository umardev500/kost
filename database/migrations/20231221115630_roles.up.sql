create table roles (
    id uuid primary key,
    name varchar(50) not null,
    tenant_id uuid,
    status status_enum default 'active'::status_enum,
    created_at timestamptz default current_timestamp,
    created_by uuid,
    updated_at timestamptz,
    updated_by uuid,
    deleted_at timestamptz,
    deleted_by uuid,
    doc_version int default 0,
    unique(tenant_id, name),
    foreign key (tenant_id)
        references tenants(id)
        on delete set null
        on update cascade
);

create index roles_id_idx on roles(id);
create index roles_tenant_id_idx on roles(tenant_id);

CREATE TRIGGER roles_update_trigger
BEFORE UPDATE ON roles
FOR EACH ROW
WHEN (
    OLD.name IS DISTINCT FROM NEW.name OR
    OLD.tenant_id IS DISTINCT FROM NEW.tenant_id OR
    OLD.status IS DISTINCT FROM NEW.status
)
EXECUTE FUNCTION update_doc_and_last_update();

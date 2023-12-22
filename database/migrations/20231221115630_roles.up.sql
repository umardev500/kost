create table roles (
    id varchar(255) primary key,
    name varchar(50) not null,
    tenant_id varchar(50),
    status status_enum default 'active'::status_enum,
    created_at timestamptz default current_timestamp,
    created_by varchar(50),
    updated_at timestamptz,
    updated_by varchar(255),
    deleted_at timestamptz,
    deleted_by varchar(255),
    doc_version int default 0,
    unique(tenant_id, name),
    foreign key (tenant_id)
        references tenants(id)
        on delete set null
        on update cascade
);

CREATE TRIGGER roles_update_trigger
BEFORE UPDATE ON roles
FOR EACH ROW
WHEN (
    OLD.name IS DISTINCT FROM NEW.name OR
    OLD.tenant_id IS DISTINCT FROM NEW.tenant_id OR
    OLD.status IS DISTINCT FROM NEW.status
)
EXECUTE FUNCTION update_doc_and_last_update();

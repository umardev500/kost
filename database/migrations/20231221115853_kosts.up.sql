create table kosts (
    id uuid primary key,
    tenant_id uuid not null,
    name varchar(255) not null,
    description text not null,
    photos jsonb not null,
    videos jsonb,
    province_id varchar(255) not null,
    city_id varchar(255) not null,
    subdistrict_id varchar(255) not null,
    village_id varchar(255) not null,
    zip_code varchar(10) not null,
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

create index kosts_id_idx on kosts(id);
create index kosts_tenant_id_idx on kosts(tenant_id);

CREATE TRIGGER kosts_update_trigger
BEFORE UPDATE ON kosts
FOR EACH ROW
WHEN (
    OLD.tenant_id IS DISTINCT FROM NEW.tenant_id OR
    OLD.name IS DISTINCT FROM NEW.name OR
    OLD.description IS DISTINCT FROM NEW.description OR
    OLD.photos IS DISTINCT FROM NEW.photos OR
    OLD.videos IS DISTINCT FROM NEW.videos OR
    OLD.province_id IS DISTINCT FROM NEW.province_id OR
    OLD.city_id IS DISTINCT FROM NEW.city_id OR
    OLD.subdistrict_id IS DISTINCT FROM NEW.subdistrict_id OR
    OLD.village_id IS DISTINCT FROM NEW.village_id OR
    OLD.zip_code IS DISTINCT FROM NEW.zip_code OR
    OLD.status IS DISTINCT FROM NEW.status
)
EXECUTE FUNCTION update_doc_and_last_update();

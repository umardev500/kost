create table tenants (
    id uuid primary key,
    name varchar(255) not null,
    id_type id_enum not null,
    id_no varchar(50) not null,
    email varchar(255) not null,
    phone varchar(12) not null,
    subdomain varchar(50) not null,
    address text not null,
    province_id varchar(255) not null,
    city_id varchar(255) not null,
    subdistrict_id varchar(255) not null,
    village_id varchar(255) not null,
    zip_code varchar(10) not null,
    status status_enum default 'inactive'::status_enum,
    created_at timestamptz default current_timestamp,
    created_by uuid,
    updated_at timestamptz,
    updated_by uuid,
    deleted_at timestamptz,
    deleted_by uuid,
    doc_version int default 0,
    unique(id_type, id_no),
    unique(email),
    unique(subdomain)
);

create index tenants_id_idx on tenants(id);

CREATE TRIGGER tenants_update_trigger
BEFORE UPDATE ON tenants
FOR EACH ROW
WHEN (
    OLD.name IS DISTINCT FROM NEW.name OR
    OLD.id_type IS DISTINCT FROM NEW.id_type OR
    OLD.id_no IS DISTINCT FROM NEW.id_no OR
    OLD.email IS DISTINCT FROM NEW.email OR
    OLD.phone IS DISTINCT FROM NEW.phone OR
    OLD.subdomain IS DISTINCT FROM NEW.subdomain OR
    OLD.address IS DISTINCT FROM NEW.address OR
    OLD.province_id IS DISTINCT FROM NEW.province_id OR
    OLD.city_id IS DISTINCT FROM NEW.city_id OR
    OLD.subdistrict_id IS DISTINCT FROM NEW.subdistrict_id OR
    OLD.village_id IS DISTINCT FROM NEW.village_id OR
    OLD.zip_code IS DISTINCT FROM NEW.zip_code OR
    OLD.status IS DISTINCT FROM NEW.status
)
EXECUTE FUNCTION update_doc_and_last_update();
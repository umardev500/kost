create table users (
    id varchar(255) primary key,
    tenant_id varchar(255), -- null indicate that user made by platform
    email varchar(50) not null,
    username varchar(50) not null,
    password varchar(255) not null,
    created_at timestamptz default current_timestamp,
    created_by varchar(50), -- null indicate that user is self registration
    updated_at timestamptz,
    updated_by varchar(255),
    deleted_at timestamptz,
    deleted_by varchar(255),
    doc_version int default 0,
    unique(email),
    unique(username),
    foreign key (tenant_id)
        references tenants(id)
        on delete set null
        on update cascade
);

CREATE TRIGGER users_update_trigger
BEFORE UPDATE ON users
FOR EACH ROW
WHEN (
    OLD.tenant_id IS DISTINCT FROM NEW.tenant_id OR
    OLD.email IS DISTINCT FROM NEW.email OR
    OLD.username IS DISTINCT FROM NEW.username OR
    OLD.password IS DISTINCT FROM NEW.password
)
EXECUTE FUNCTION update_doc_and_last_update();

create table if not exists inquilino
(
    id         int auto_increment
    primary key,
    dni        int         not null,
    nombre     varchar(40) not null,
    apellido   varchar(40) not null,
    direccion  varchar(50) not null,
    telefono   int         not null,
    deleted_at date        null,
    updated_at date        null,
    created_at date        null
    );

create table if not exists propietarios
(
    id         int auto_increment
    primary key,
    dni        int          not null,
    apellido   varchar(40)  not null,
    nombre     varchar(40)  not null,
    telefono   int          not null,
    mail       varchar(320) null,
    password   varchar(10)  not null,
    deleted_at date         null,
    created_at date         null,
    updated_at date         null
    )
    auto_increment = 4;

create table if not exists inmuebles
(
    id             int auto_increment
    primary key,
    direccion      varchar(50) not null,
    ambientes      int         not null,
    tipo           varchar(15) not null,
    uso            varchar(15) not null,
    precio         double      not null,
    disponible     tinyint(1)  not null,
    propietario_id int         not null,
    deleted_at     datetime    null,
    updated_at     datetime    null,
    created_at     datetime    null,
    constraint inmuebles_ibfk_1
    foreign key (propietario_id) references propietarios (id)
    )
    auto_increment = 7;

create table if not exists alquiler
(
    id           int auto_increment
    primary key,
    precio       double   not null,
    fecha_inicio datetime not null,
    fecha_fin    datetime not null,
    inquilino_id int      not null,
    inmueble_id  int      not null,
    deleted_at   date     null,
    updated_at   date     null,
    created_at   date     null,
    constraint alquiler_inmueble_inmuebleID_fk
    foreign key (inmueble_id) references inmuebles (id),
    constraint alquiler_inquilino_inquilinoID_fk
    foreign key (inquilino_id) references inquilino (id)
    );

create index propietario
    on inmuebles (propietario_id);

create table if not exists pagos
(
    id          int auto_increment
    primary key,
    numero_pago int      not null,
    alquiler_id int      not null,
    fecha       datetime not null,
    importe     double   not null,
    deleted_at  date     null,
    created_at  date     null,
    updated_at  date     null,
    constraint pagos_alquiler_alquilerID_fk
    foreign key (alquiler_id) references alquiler (id)
    );
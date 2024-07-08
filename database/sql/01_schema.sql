CREATE TABLE usuario (
    usuario_id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    telefono VARCHAR(15),
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    activo BOOLEAN DEFAULT FALSE
);

CREATE TABLE comunidad (
    comunidad_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    direccion_calle VARCHAR(100),
    direccion_numero VARCHAR(20),
    direccion_colonia VARCHAR(100),
    direccion_cp VARCHAR(10),
    direccion_ciudad VARCHAR(100),
    direccion_estado VARCHAR(100),
    direccion_pais VARCHAR(100),
    tipo VARCHAR(50) NOT NULL CHECK (tipo IN ('Fraccionamiento', 'Edificio', 'Calle')),
    modelo_suscripcion VARCHAR(20) NOT NULL CHECK (modelo_suscripcion IN ('Mensual', 'Bimestral', 'Anual'))
);

CREATE TABLE suscripcion (
     suscripcion_id SERIAL PRIMARY KEY,
     usuario_id INT REFERENCES usuario(usuario_id),
     comunidad_id INT REFERENCES comunidad(comunidad_id),
     modelo_suscripcion VARCHAR(20) NOT NULL CHECK (modelo_suscripcion IN ('Mensual', 'Bimestral', 'Anual')),
     fecha_inicio DATE NOT NULL,
     fecha_fin DATE NOT NULL,
     monto DECIMAL(10, 2) NOT NULL
);

CREATE TABLE pago (
  pago_id SERIAL PRIMARY KEY,
  suscripcion_id INT REFERENCES suscripcion(suscripcion_id),
  fecha_pago DATE NOT NULL,
  monto DECIMAL(10, 2) NOT NULL
);

CREATE TABLE casa (
    casa_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    direccion VARCHAR(255) NOT NULL,
    numero INT NOT NULL
);

CREATE TABLE departamento (
    departamento_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    direccion VARCHAR(255) NOT NULL,
    numero INT NOT NULL
);

CREATE TABLE habitante (
    habitante_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    telefono VARCHAR(15),
    email VARCHAR(100)
);

CREATE TABLE comite (
    comite_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE comite_miembro (
    comite_miembro_id SERIAL PRIMARY KEY,
    comite_id INT REFERENCES comite(comite_id),
    habitante_id INT REFERENCES habitante(habitante_id)
);

CREATE TABLE junta (
    junta_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    fecha DATE NOT NULL,
    descripcion TEXT
);

CREATE TABLE anuncio (
    anuncio_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    fecha DATE NOT NULL,
    descripcion TEXT
);

CREATE TABLE cuota (
    cuota_id SERIAL PRIMARY KEY,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    monto DECIMAL(10, 2) NOT NULL,
    fecha_pago DATE NOT NULL,
    descripcion TEXT
);

CREATE TABLE bazar (
    bazar_id SERIAL PRIMARY KEY,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    fecha DATE NOT NULL,
    descripcion TEXT,
    precio DECIMAL(10, 2) NOT NULL,
    vendido BOOLEAN DEFAULT FALSE,
    acepta_regateo BOOLEAN DEFAULT FALSE
);

CREATE TABLE anuncio_casa (
    anuncio_casa_id SERIAL PRIMARY KEY,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    fecha DATE NOT NULL,
    descripcion TEXT
);

CREATE TABLE anuncio_comite (
     anuncio_id SERIAL PRIMARY KEY,
     comite_id INT REFERENCES comite(comite_id),
     fecha DATE NOT NULL,
     descripcion TEXT NOT NULL
);

-- Indexes:
CREATE INDEX idx_usuario_email ON usuario(email);
CREATE INDEX idx_comunidad_nombre ON comunidad(nombre);
CREATE INDEX idx_suscripcion_usuario_id ON suscripcion(usuario_id);
CREATE INDEX idx_suscripcion_comunidad_id ON suscripcion(comunidad_id);
CREATE INDEX idx_casa_comunidad_id ON casa(comunidad_id);
CREATE INDEX idx_departamento_comunidad_id ON departamento(comunidad_id);
CREATE INDEX idx_habitante_casa_id ON habitante(casa_id);
CREATE INDEX idx_habitante_departamento_id ON habitante(departamento_id);
CREATE INDEX idx_comite_comunidad_id ON comite(comunidad_id);
CREATE INDEX idx_comite_miembro_comite_id ON comite_miembro(comite_id);
CREATE INDEX idx_comite_miembro_habitante_id ON comite_miembro(habitante_id);
CREATE INDEX idx_junta_comunidad_id ON junta(comunidad_id);
CREATE INDEX idx_anuncio_comunidad_id ON anuncio(comunidad_id);
CREATE INDEX idx_anuncio_casa_casa_id ON anuncio_casa(casa_id);
CREATE INDEX idx_anuncio_casa_departamento_id ON anuncio_casa(departamento_id);
CREATE INDEX idx_bazar_casa_id ON bazar(casa_id);
CREATE INDEX idx_bazar_departamento_id ON bazar(departamento_id);

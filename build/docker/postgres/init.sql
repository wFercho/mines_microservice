CREATE EXTENSION postgis;

CREATE TABLE minas (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    empresa_asociada VARCHAR(100),
    coordenadas GEOGRAPHY(POINT, 4326),
    departamento VARCHAR(100),
    municipio VARCHAR(100),
    barrio_vereda VARCHAR(100),
    direccion VARCHAR(255),
    estado_operativo VARCHAR(50),
    mqtt_topics TEXT[]
);

INSERT INTO minas (nombre, empresa_asociada, coordenadas, departamento, municipio, barrio_vereda, direccion, estado_operativo, mqtt_topics) VALUES
('Mina El Diamante', 'Carbones de Boyacá S.A.', ST_GeogFromText('POINT(-73.3675 5.535)', 4326), 'Boyacá', 'Sogamoso', 'Vereda Morcá', 'Km 5 vía Morcá', 'Activa', ARRAY[]),
('Mina La Esperanza', 'Minerales del Centro Ltda.', ST_GeogFromText('POINT(-73.214 5.583)', 4326), 'Boyacá', 'Nobsa', 'Vereda Las Caleras', 'Sector Las Caleras', 'Activa', ARRAY[]),
('Mina Santa Teresa', 'C.I. Milpa S.A.', ST_GeogFromText('POINT(-73.332 5.715)', 4326), 'Boyacá', 'Paipa', 'Vereda El Salitre', 'Finca El Salitre', 'Inactiva', ARRAY[]),
('Mina El Porvenir', 'Cooperativa Minera de Samacá', ST_GeogFromText('POINT(-73.471 5.492)', 4326), 'Boyacá', 'Samacá', 'Vereda San Roque', 'Vía San Roque', 'Activa', ARRAY[]),
('Mina La Fortuna', 'Boyacá Coal Ltd.', ST_GeogFromText('POINT(-73.273 5.454)', 4326), 'Boyacá', 'Duitama', 'Vereda La Pradera', 'Calle 10 #15-20', 'Activa', ARRAY[]),
('Mina San José', 'Compañía Minera del Centro', ST_GeogFromText('POINT(-73.356 5.576)', 4326), 'Boyacá', 'Tibasosa', 'Vereda El Palmar', 'Carrera 8 #12-30', 'Cerrada', ARRAY[]),
('Mina La Cumbre', 'Mineros de Nobsa SAS', ST_GeogFromText('POINT(-73.233 5.583)', 4326), 'Boyacá', 'Nobsa', 'Vereda El Espino', 'Km 2 vía El Espino', 'Activa', ARRAY[]),
('Mina El Tesoro', 'Carbonífera Andina', ST_GeogFromText('POINT(-73.415 5.555)', 4326), 'Boyacá', 'Firavitoba', 'Vereda El Hato', 'Sector El Hato', 'Activa', ARRAY[]),
('Mina La Gloria', 'Carbones del Oriente', ST_GeogFromText('POINT(-73.292 5.617)', 4326), 'Boyacá', 'Belén', 'Vereda La Lajita', 'Finca La Lajita', 'Inactiva', ARRAY[]),
('Mina San Antonio', 'Minerales de Occidente', ST_GeogFromText('POINT(-73.481 5.721)', 4326), 'Boyacá', 'Cómbita', 'Vereda San Isidro', 'Vía San Isidro', 'Activa', ARRAY[]),
('Mina El Milagro', 'C.I. Colombiana de Carbón', ST_GeogFromText('POINT(-73.361 5.432)', 4326), 'Boyacá', 'Corrales', 'Vereda La Esmeralda', 'Carrera 5 #10-15', 'Activa', ARRAY[]),
('Mina La Victoria', 'Carbones del Alto Chicamocha', ST_GeogFromText('POINT(-73.257 5.768)', 4326), 'Boyacá', 'Tutazá', 'Vereda El Diamante', 'Km 3 vía El Diamante', 'Cerrada', ARRAY[]),
('Mina Santa Ana', 'Minería Integral SAS', ST_GeogFromText('POINT(-73.412 5.612)', 4326), 'Boyacá', 'Toca', 'Vereda La Laguna', 'Sector La Laguna', 'Activa', ARRAY[]),
('Mina El Edén', 'Carbones y Coques Boyacá', ST_GeogFromText('POINT(-73.295 5.492)', 4326), 'Boyacá', 'Busbanzá', 'Vereda El Paraíso', 'Calle 2 #3-25', 'Inactiva', ARRAY[]),
('Mina La Unión', 'Cooperativa Minera de Tópaga', ST_GeogFromText('POINT(-73.298 5.748)', 4326), 'Boyacá', 'Tópaga', 'Vereda La Chapa', 'Vía La Chapa', 'Activa', ARRAY[]),
('Mina San Miguel', 'Minerales de Boyacá Ltda.', ST_GeogFromText('POINT(-73.341 5.678)', 4326), 'Boyacá', 'Monguí', 'Vereda El Carmen', 'Carrera 7 #8-40', 'Activa', ARRAY[]),
('Mina La Prosperidad', 'Carbones del Suroccidente', ST_GeogFromText('POINT(-73.219 5.832)', 4326), 'Boyacá', 'Gámeza', 'Vereda La Vega', 'Km 1 vía La Vega', 'Inactiva', ARRAY[]),
('Mina El Refugio', 'C.I. Minas Paz de Río', ST_GeogFromText('POINT(-73.271 5.791)', 4326), 'Boyacá', 'Paz de Río', 'Vereda El Refugio', 'Sector El Refugio', 'Activa', ARRAY[]),
('Mina La Ilusión', 'Mineros de Socotá SAS', ST_GeogFromText('POINT(-73.315 5.946)', 4326), 'Boyacá', 'Socotá', 'Vereda El Mirador', 'Finca El Mirador', 'Cerrada', ARRAY[]),
('Mina San Rafael', 'Carbones de la Uvita', ST_GeogFromText('POINT(-73.026 6.317)', 4326), 'Boyacá', 'La Uvita', 'Vereda San Rafael', 'Km 4 vía San Rafael', 'Activa', ARRAY[]);
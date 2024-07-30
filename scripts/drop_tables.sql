\echo -n '¿Estás seguro de querer eliminar las tablas? (yes/y/s/sí): '
\prompt confirmacion

SET my.confirmacion = :'confirmacion';

DO $do$
    DECLARE
        respuesta TEXT;
    BEGIN
        SELECT current_setting('my.confirmacion') INTO respuesta;

        IF respuesta IS NULL OR respuesta = '' THEN
            RAISE EXCEPTION 'No se proporcionó una respuesta.';
        END IF;

        respuesta := lower(respuesta);

        IF respuesta IN ('yes', 'y', 's', 'sí', 'si') THEN
            EXECUTE 'DROP TABLE IF EXISTS bazar CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS cuota CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS anuncio_casa CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS anuncio_comite CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS anuncio CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS junta CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS comite_miembro CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS comite CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS habitante CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS departamento CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS casa CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS pago CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS suscripcion CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS comunidad CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS confirmacion_cuenta CASCADE';
            EXECUTE 'DROP TABLE IF EXISTS usuario CASCADE';

            RAISE NOTICE 'Las tablas han sido droppeadas.';
        ELSIF respuesta = 'no' THEN
            RAISE NOTICE 'Operación cancelada. Las tablas no han sido modificadas.';
        ELSE
            RAISE EXCEPTION 'Respuesta no válida. Por favor, responde "yes", "y", "s", "sí" o "no".';
        END IF;
    END
$do$;

RESET my.confirmacion;

\echo 'Script ejecutado.'
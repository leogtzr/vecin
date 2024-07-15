\echo -n '¿Estás seguro de querer eliminar las tablas? (yes/no): '
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

        CASE lower(respuesta)
            WHEN 'yes' THEN
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
            WHEN 'no' THEN
                RAISE NOTICE 'Operación cancelada. La tabla no ha sido modificada.';
            ELSE
                RAISE EXCEPTION 'Respuesta no válida. Por favor, responde "yes" o "no".';
            END CASE;
    END
$do$;

RESET my.confirmacion;

\echo 'Script ejecutado.'

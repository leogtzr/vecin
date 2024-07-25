$(document).ready(function() {
    let selectedFraccionamientoId = null;

    // Cargar lista de fraccionamientos
    function loadFraccionamientos() {
        $.ajax({
            url: '/api/fraccionamientos',
            type: 'GET',
            success: function(response) {
                let tableBody = $('#fraccionamientosTable tbody');
                tableBody.empty();
                console.log(response);
                response.forEach(function(fraccionamiento) {
                    console.log('debug:x begin');
                    console.log(fraccionamiento);
                    console.log('debug:x end');
                    tableBody.append(`
                        <tr data-id="${fraccionamiento.comunidad_id}">
                            <td>${fraccionamiento.name}</td>
                            <td>${fraccionamiento.tipo}</td>
                            <td>${fraccionamiento.direccion_estado}</td>
                            <td>${fraccionamiento.direccion_ciudad}</td>
                            <td>
                                <button class="btn btn-sm btn-info edit-btn">Editar</button>
                                <button class="btn btn-sm btn-danger delete-btn">Eliminar</button>
                            </td>
                        </tr>
                    `);
                });
            },
            error: function(xhr, status, error) {
                alert('Error al cargar fraccionamientos: ' + error);
            }
        });
    }

    // Cargar detalles de fraccionamiento
    function loadFraccionamientoDetails(id) {
        $.ajax({
            url: `/api/fraccionamientos/${id}`,
            type: 'GET',
            success: function(fraccionamiento) {
                $('#nombreFraccionamiento').val(fraccionamiento.nombre);
                $('#tipoFraccionamiento').val(fraccionamiento.tipo);
                $('#modeloSuscripcion').val(fraccionamiento.modelo_suscripcion);
                $('#direccionCalle').val(fraccionamiento.direccion_calle);
                $('#direccionNumero').val(fraccionamiento.direccion_numero);
                $('#direccionColonia').val(fraccionamiento.direccion_colonia);
                $('#direccionCP').val(fraccionamiento.direccion_cp);
                $('#direccionCiudad').val(fraccionamiento.direccion_ciudad);

                // Habilitar edición
                $('.fraccionamiento-details input').prop('readonly', false);
                $('#saveFraccionamiento').show();
            },
            error: function(xhr, status, error) {
                alert('Error al cargar detalles del fraccionamiento: ' + error);
            }
        });
    }

    // Evento para editar fraccionamiento
    $(document).on('click', '.edit-btn', function() {
        selectedFraccionamientoId = $(this).closest('tr').data('id');
        loadFraccionamientoDetails(selectedFraccionamientoId);
    });

    // Guardar cambios del fraccionamiento
    $('#saveFraccionamiento').on('click', function() {
        let fraccionamientoData = {
            nombre: $('#nombreFraccionamiento').val(),
            tipo: $('#tipoFraccionamiento').val(),
            modelo_suscripcion: $('#modeloSuscripcion').val(),
            direccion_calle: $('#direccionCalle').val(),
            direccion_numero: $('#direccionNumero').val(),
            direccion_colonia: $('#direccionColonia').val(),
            direccion_cp: $('#direccionCP').val(),
            direccion_ciudad: $('#direccionCiudad').val()
        };

        $.ajax({
            url: `/api/fraccionamientos/${selectedFraccionamientoId}`,
            type: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(fraccionamientoData),
            success: function(response) {
                alert('Fraccionamiento actualizado con éxito');
                loadFraccionamientos();
                $('.fraccionamiento-details input').prop('readonly', true);
                $('#saveFraccionamiento').hide();
            },
            error: function(xhr, status, error) {
                alert('Error al actualizar fraccionamiento: ' + error);
            }
        });
    });

    // Eliminar fraccionamiento
    $(document).on('click', '.delete-btn', function() {
        if (confirm('¿Está seguro de que desea eliminar este fraccionamiento?')) {
            let fraccionamientoId = $(this).closest('tr').data('id');
            $.ajax({
                url: `/api/fraccionamientos/${fraccionamientoId}`,
                type: 'DELETE',
                success: function(response) {
                    alert('Fraccionamiento eliminado con éxito');
                    loadFraccionamientos();
                },
                error: function(xhr, status, error) {
                    alert('Error al eliminar fraccionamiento: ' + error);
                }
            });
        }
    });

    // Cargar fraccionamientos al iniciar
    loadFraccionamientos();
});
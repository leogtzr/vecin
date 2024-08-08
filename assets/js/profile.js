import { showModal } from './modalHelper.js';

$(document).ready(function() {
    function isValidUsername(username) {
        const trimmedUsername = username.trim();
        const usernameRegex = /^[a-zA-Z0-9]{5,14}$/;

        return usernameRegex.test(trimmedUsername);
    }

    $('#perfilForm').on('submit', function(e) {
        e.preventDefault();

        console.log('debug:x click');
    });

    //let selectedFraccionamientoId = null;

    // Cargar lista de fraccionamientos
    // function loadFraccionamientos() {
    //     $.ajax({
    //         url: '/api/fraccionamientos',
    //         type: 'GET',
    //         success: function(response) {
    //             let tableBody = $('#fraccionamientosTable tbody');
    //             tableBody.empty();
    //             console.log(response);
    //             response.forEach(function(fraccionamiento) {
    //                 console.log('debug:x begin');
    //                 console.log(fraccionamiento);
    //                 console.log('debug:x end');
    //                 tableBody.append(`
    //                     <tr data-id="${fraccionamiento.community_id}">
    //                         <td>${fraccionamiento.name}</td>
    //                         <td>${fraccionamiento.tipo}</td>
    //                         <td>${fraccionamiento.direccion_estado}</td>
    //                         <td>${fraccionamiento.direccion_ciudad}</td>
    //                         <td>
    //                             <button class="btn btn-sm btn-info edit-btn">Editar ✏️</button>
    //                             <button class="btn btn-sm btn-danger delete-btn" disabled>Eliminar 🗑</button>
    //                         </td>
    //                     </tr>
    //                 `);
    //             });
    //         },
    //         error: function(xhr, status, error) {
    //             alert('Error al cargar fraccionamientos: ' + error);
    //         }
    //     });
    // }

    function enableComponentsForEdition() {
        $('#nombreFraccionamiento').prop('readonly', false);
        $('#modeloSuscripcion').prop('readonly', false);
        $('#direccionColonia').prop('readonly', false);
        $('#direccionCP').prop('readonly', false);
        $('#referencias').prop('readonly', false);
        $('#descripcion').prop('readonly', false);
        $('#direccionCalle').prop('readonly', false);
        // TODO: might need to add more elements here...
    }

    function disableComponentsForEdition() {
        $('#nombreFraccionamiento').prop('readonly', true);
        $('#modeloSuscripcion').prop('readonly', true);
        $('#direccionColonia').prop('readonly', true);
        $('#direccionCP').prop('readonly', true);
        $('#referencias').prop('readonly', true);
        $('#descripcion').prop('readonly', true);
        $('#direccionCalle').prop('readonly', true);
        // TODO: might need to add more elements here...
    }

    // Cargar detalles de fraccionamiento
    // function loadFraccionamientoDetails(id) {
    //     $.ajax({
    //         url: `/api/fraccionamientos/${id}`,
    //         type: 'GET',
    //         success: function(fraccionamiento) {
    //             console.log('debug:x got:');
    //             console.log(fraccionamiento);
    //
    //             var modeloSuscripcion = $('#modeloSuscripcion');
    //
    //             modeloSuscripcion.empty();
    //             const opciones = ["Mensual", "Anual"];
    //             opciones.forEach(function(opcion) {
    //                 const optionElement = $('<option>', {
    //                     value: opcion,
    //                     text: opcion
    //                 });
    //                 modeloSuscripcion.append(optionElement);
    //             });
    //
    //             modeloSuscripcion.val(fraccionamiento.modelo_suscripcion);
    //
    //             $('#nombreFraccionamiento').val(fraccionamiento.name);
    //             $('#tipoFraccionamiento').val(fraccionamiento.tipo);
    //             $('#direccionCalle').val(fraccionamiento.direccion_calle);
    //             $('#direccionNumero').val(fraccionamiento.direccion_numero);
    //             $('#direccionColonia').val(fraccionamiento.direccion_colonia);
    //             $('#direccionCP').val(fraccionamiento.direccion_cp);
    //             $('#direccionCiudad').val(fraccionamiento.direccion_ciudad);
    //             $('#direccionEstado').val(fraccionamiento.direccion_estado);
    //             $('#referencias').val(fraccionamiento.referencias);
    //             $('#descripcion').val(fraccionamiento.descripcion);
    //
    //             // Habilitar edición:
    //             enableComponentsForEdition();
    //
    //             $('#saveFraccionamiento').show();
    //             $('#cancelSaveFraccionamiento').show();
    //         },
    //         error: function(xhr, status, error) {
    //             // TODO: improve this:
    //             alert('Error al cargar detalles del fraccionamiento: ' + error);
    //         }
    //     });
    // }

    // Evento para editar fraccionamiento
    // $(document).on('click', '.edit-btn', function() {
    //     const row = $(this).closest('tr');
    //     selectedFraccionamientoId = row.data('id');
    //     console.log('Selected row data-id: ' + row.data('id'));
    //     console.log('Selected fraccionamiento ID: ' + selectedFraccionamientoId);
    //
    //     loadFraccionamientoDetails(selectedFraccionamientoId);
    // });

    // $('#cancelSaveFraccionamiento').on('click', function() {
    //     disableComponentsForEdition();
    //     $('#saveFraccionamiento').hide();
    //     $('#cancelSaveFraccionamiento').hide();
    // });

    // Guardar cambios del fraccionamiento
    // $('#saveFraccionamiento').on('click', function() {
    //     let fraccionamientoData = {
    //         nombreComunidad: $('#nombreFraccionamiento').val(),
    //         tipoComunidad: $('#tipoFraccionamiento').val(),
    //         modeloSuscripcion: $('#modeloSuscripcion').val(),
    //         direccionCalle: $('#direccionCalle').val(),
    //         direccionNumero: $('#direccionNumero').val(),
    //         direccionColonia: $('#direccionColonia').val(),
    //         direccionCodigoPostal: $('#direccionCP').val(),
    //         direccionCiudad: $('#direccionCiudad').val(),
    //         direccionEstado: $('#direccionEstado').val(),
    //         referencias: $('#referencias').val(),
    //         descripcion: $('#descripcion').val(),
    //     };
    //
    //     console.log('debug:x Sending to update:');
    //     console.log(fraccionamientoData);
    //
    //     $.ajax({
    //         url: `/api/fraccionamientos/${selectedFraccionamientoId}`,
    //         type: 'PUT',
    //         contentType: 'application/json',
    //         data: JSON.stringify(fraccionamientoData),
    //         success: function(response) {
    //             console.log(response);
    //             //alert('Fraccionamiento actualizado con éxito');
    //             loadFraccionamientos();
    //             disableComponentsForEdition();
    //             $('#saveFraccionamiento').hide();
    //             $('#cancelSaveFraccionamiento').hide();
    //
    //             showModal('Fraccionamiento actualizado con éxito', 'success');
    //         },
    //         error: function(xhr, status, error) {
    //             // TODO: improve this:
    //             //alert('Error al actualizar fraccionamiento: ' + error);
    //             console.log(status);
    //             console.log(error);
    //             console.log(xhr);
    //             // showErrorModal('Error al actualizar fraccionamiento: ' + error);
    //             showModal('Error al cargar fraccionamientos: ' + error, 'error');
    //         }
    //     });
    // });

    // Eliminar fraccionamiento
    // $(document).on('click', '.delete-btn', function() {
    //     if (confirm('¿Está seguro de que desea eliminar este fraccionamiento?')) {
    //         let fraccionamientoId = $(this).closest('tr').data('id');
    //         $.ajax({
    //             url: `/api/fraccionamientos/${fraccionamientoId}`,
    //             type: 'DELETE',
    //             success: function(response) {
    //                 alert('Fraccionamiento eliminado con éxito');
    //                 loadFraccionamientos();
    //             },
    //             error: function(xhr, status, error) {
    //                 // TODO: improve this:
    //                 alert('Error al eliminar fraccionamiento: ' + error);
    //             }
    //         });
    //     }
    // });

    // Cargar fraccionamientos al iniciar
    // loadFraccionamientos();
});
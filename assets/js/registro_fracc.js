$(document).ready(function() {
    function resetCiudadField() {
        var $direccionCiudad = $('#direccionCiudad');
        $direccionCiudad.empty();
        $direccionCiudad.append('<option value="">Selecciona una ciudad</option>');
    }

    function resetEstadoField() {
        var $direccionEstado = $('#direccionEstado');
        $direccionEstado.empty();
        $direccionEstado.append('<option value="">Selecciona un estado</option>');
    }

    function fetchStatesByCountry(countryId) {
        var url = `api/region?geonameId=${countryId}`;

        $.getJSON(url, function(data) {
            var estados = data.geonames;
            var $direccionEstado = $('#direccionEstado');

            $direccionEstado.empty();
            $direccionEstado.append('<option value="">Selecciona un estado</option>');

            estados.forEach(function(estado) {
                var option =
                    `<option value="${estado.name}" name="${estado.name}" data-regionId="${estado.geonameId}">${estado.name}</option>`;
                $direccionEstado.append(option);
            });

            $direccionEstado.trigger('change');
        });
    }

    // cities or municipalities...
    function fetchCitiesByStateId(stateId) {
        var ciudadesUrl = `api/region?geonameId=${stateId}`;

        $.getJSON(ciudadesUrl, function(data) {
            var ciudades = data.geonames;
            var $direccionCiudad = $('#direccionCiudad');

            $direccionCiudad.empty();
            $direccionCiudad.append('<option value="">Selecciona una ciudad</option>');

            ciudades.forEach(function(ciudad) {
                var option = `<option value="${ciudad.toponymName}" name="${ciudad.toponymName}">${ciudad.toponymName}</option>`;
                $direccionCiudad.append(option);
            });
        }).fail(function() {
            // ToDo: improve error handling here.
            console.log('Error calling URL');
        });
    }

        $('#registroComunidadForm').on('submit', function(e) {
            e.preventDefault();

            var formData = new FormData(this);
            const clickedElement = $(this);

            var formData = {
                nombreComunidad: $('#nombreComunidad').val(),
                tipoComunidad: $('#tipoComunidad').val(),
                modeloSuscripcion: $('#modeloSuscripcion').val(),
                direccionCalle: $('#direccionCalle').val(),
                direccionNumero: $('#direccionNumero').val(),
                direccionColonia: $('#direccionColonia').val(),
                direccionCodigoPostal: $('#direccionCodigoPostal').val(),
                direccionEstado: $('#direccionEstado').val(),
                direccionCiudad: $('#direccionCiudad').val(),
                direccionPais: $('#direccionPais').val(),
                referencias: $('#referencias').val(),
                descripcion: $('#descripcion').val(),
                registranteNombre: $('#registranteNombre').val(),
                registranteApellido: $('#registranteApellido').val(),
                registranteTelefono: $('#registranteTelefono').val(),
                registranteUsername: $('#registranteUsername').val(),
                registranteEmail: $('#registranteEmail').val(),
                habitante: $('#habitante').val(),
                registranteSignUpUserName: $('#registranteSignUpUserName').val(),
                registranteSignUpPassword: $('#registranteSignUpPassword').val()
            };

            console.log('Will try to send: ', formData);
            $.ajax({
                url: '/registrar-fracc',
                type: 'POST',
                data: JSON.stringify(formData),
                contentType: 'application/json',
                success: function(response) {
                    console.log('OK');
                    console.log(response);

                    const infoModal = clickedElement.find('.info-modal');
                    $('#infoModal').modal('show');

                },
                error: function(xhr, status, error) {
                    console.error('Error:', error);
                }
            });
        });

    $('#direccionEstado').on('change', function() {
        var estadoElement = $('#direccionEstado option:selected');
        var estado = estadoElement.attr('data-regionId');

        var pais = $('#direccionPais').val();

        // TODO: quizás no necesitamos verificar el país.
        if (estado && pais) {
            fetchCitiesByStateId(estado);
        }
    });

    $('#direccionPais').on('change', function() {
        const pais = $('#direccionPais').val();
        if (pais) {
            console.log('debug:x there was a change in the pais:', pais);

            resetEstadoField();
            resetCiudadField();

            fetchStatesByCountry(pais);
        }
    }).trigger('change');

    if ($("#registroComunidadForm").length > 0) {
        var geonameIdMexico = 3996063; // default: ID de México

        // ToDo: change this to our WS endpoint:
        fetchStatesByCountry(geonameIdMexico);
    }

    $('#infoModal').on('click', '#okButton', function() {
        window.location.href = '/view-fraccs';
    });

});

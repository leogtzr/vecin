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
            console.log("nombreComunidad=", $('#nombreComunidad').val());

            const clickedElement = $(this);

            for (var pair of formData.entries()) {
                console.log(pair[0] + ': ' + pair[1]);
            }

            var formData = {};
            formData['nombreComunidad'] = $('#nombreComunidad').val();
            formData['tipoComunidad'] = $('#tipoComunidad').val();
            formData['modeloSuscripcion'] = $('#modeloSuscripcion').val();
            formData['direccionCalle'] = $('#direccionCalle').val();
            formData['direccionNumero'] = $('#direccionNumero').val();
            formData['direccionColonia'] = $('#direccionColonia').val();
            formData['direccionCodigoPostal'] = $('#direccionCodigoPostal').val();
            formData['direccionEstado'] = $('#direccionEstado').val();
            formData['direccionCiudad'] = $('#direccionCiudad').val();
            formData['direccionPais'] = $('#direccionPais').val();
            formData['referencias'] = $('#referencias').val();
            formData['descripcion'] = $('#descripcion').val();
            formData['registranteNombre'] = $('#registranteNombre').val();
            formData['registranteApellido'] = $('#registranteApellido').val();
            formData['registranteTelefono'] = $('#registranteTelefono').val();
            formData['registranteEmail'] = $('#registranteEmail').val();
            formData['habitante'] = $('#habitante').val();
            formData['registranteSignUpUserName'] = $('#registranteSignUpUserName').val();
            formData['registranteSignUpPassword'] = $('#registranteSignUpPassword').val();

            console.log('Will try to send: ', formData);
            $.ajax({
                url: '/registrar-fracc',
                type: 'POST',
                data: JSON.stringify(formData),
                contentType: 'application/json',
                success: function(response) {
                    console.log('OK', response);
                    const infoModal = clickedElement.find('.info-modal');
                    infoModal.text('Fraccionamiento registrado correctamente');
                    infoModal.show();

                    setTimeout(() => {
                        infoModal.hide();
                        // window.location.href = `modify?book_id=${bookID}`;
                    }, 800);
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

        console.log('debug:x Change event for estado:', estado, ', pais:', pais);
        console.log('debug:x estado', estadoElement.val());

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

});

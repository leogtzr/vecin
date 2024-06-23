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
        var username = 'leogtzr';

        var url = `api/region?geonameId=${countryId}&username=${username}`;

        $.getJSON(url, function(data) {
            var estados = data.geonames;
            var $direccionEstado = $('#direccionEstado');
            
            $direccionEstado.empty();
            $direccionEstado.append('<option value="">Selecciona un estado</option>');
            
            estados.forEach(function(estado) {
                var option = `<option value="${estado.geonameId}">${estado.name}</option>`;
                $direccionEstado.append(option);
            });
        });
    }

    // cities or municipalities...
    function fetchCitiesByStateId(stateId) {
        var username = 'leogtzr';
        var ciudadesUrl = `api/region?geonameId=${stateId}&username=${username}`;

        $.getJSON(ciudadesUrl, function(data) {
            var ciudades = data.geonames;
            var $direccionCiudad = $('#direccionCiudad');

            $direccionCiudad.empty();
            $direccionCiudad.append('<option value="">Selecciona una ciudad</option>');

            ciudades.forEach(function(ciudad) {
                var option = `<option value="${ciudad.geonameId}">${ciudad.toponymName}</option>`;
                $direccionCiudad.append(option);
            });
        }).fail(function() {
            // ToDo: improve error handling here.
            console.log('Error calling URL');
        });
    }

    $('#direccionEstado').on('change', function() {
        const estado = $('#direccionEstado').val();
        const pais = $('#direccionPais').val();
        
        if (estado && pais) {
            fetchCitiesByStateId(estado);
        }
    }).trigger('change');

    $('#direccionPais').change('change', function() {
        const pais = $('#direccionPais').val();
        if (pais) {
            console.log('debug:x there was a change in the pais:', pais);

            resetEstadoField();
            resetCiudadField();

            fetchStatesByCountry(pais);
        }
    }).trigger('change');

    if ($("#registroComunidadForm").length > 0) {
        var geonameIdMexico = 3996063; // ID de México
        // ToDo: put this behind an environmental variable.
        
        // ToDo: change this to our WS endpoint:
        fetchStatesByCountry(geonameIdMexico);
    }
});
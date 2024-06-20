$(document).ready(function() {
    // Use mouseenter and mouseleave events to show and hide the modal
    // $('.card').on('mouseenter', function() {
    //     console.log('I have entered');
    //     var targetModal = $(this).data('target');
    //     $(targetModal).modal('show');
    // }).on('mouseleave', function() {
    //     console.log('I have exit');
    //     var targetModal = $(this).data('target');
    //     $(targetModal).modal('hide');
    // });

    $('#tipoComunidad').on('change', function() {
        var tipoComunidad = $('#tipoComunidad').val();

        console.log('It has changed to: ' + tipoComunidad);

        $('#tipoComunidadSpan').text(tipoComunidad.toLowerCase());
    }).trigger('change');

    $('#habitante').on('change', function() {
        if ($(this).val() === 'Si') {
            console.log('YES Opening this...');
            $('#registranteSignUp').slideDown();
            $('#registranteSignUpUserName, #registranteSignUpPassword').attr('required', true);
        } else {
            console.log('NO ...');
            $('#registranteSignUp').slideUp();
            $('#registranteSignUpUserName, #registranteSignUpPassword').removeAttr('required');
        }
    });

    $('#direccionEstado').on('change', function() {
        const estado = $('#direccionEstado').val();
        const pais = $('#direccionPais').val();

        if (estado && pais) {
            console.log('1) Buscar para:', estado, "y", pais);
            // fetchCities(pais, estado);
        } else {
            console.log('1) No Buscar para:', estado, "y", pais);
        }
    }).trigger('change');

    $('#direccionPais').change('change', function() {
        const estado = $('#direccionEstado').val();
        const pais = $('#direccionPais').val();

        if (estado && pais) {
            //fetchCities(pais, estado);
            console.log('2) Buscar para:', estado, "y", pais);
        } else {
            console.log('2) No Buscar para:', estado, "y", pais);
        }
    }).trigger('change');
});
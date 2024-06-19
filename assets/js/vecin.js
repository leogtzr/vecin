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
        $('#tipoComunidadSpan').text(tipoComunidad.toLowerCase());
    }).trigger('change');

    $('#habitante').on('change', function() {
        if ($(this).val() === 'Si') {
            console.log('YES Opening this...');
            $('#usuarioContraseña').slideDown();
            $('#nombreUsuario, #contrasena').attr('required', true);
        } else {
            console.log('NO ...');
            $('#usuarioContraseña').slideUp();
            $('#nombreUsuario, #contrasena').removeAttr('required');
        }
    });
});
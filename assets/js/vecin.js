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
        if ($(this).val() === 'yes') {
            console.log('YES Opening this...');
            $('#registranteSignUp').slideDown();
            $('#registranteSignUpUserName, #registranteSignUpPassword').attr('required', true);
        } else {
            $('#registranteSignUp').slideUp();
            $('#registranteSignUpUserName, #registranteSignUpPassword').removeAttr('required');
        }
    });
});
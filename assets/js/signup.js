$(document).ready(function() {
    $('#signUp').on('submit', function(e) {
        e.preventDefault();

        var password = $('#password').val();
        var confirmPassword = $('#confirm_password').val();

        if (password !== confirmPassword) {
            $('#alert').fadeIn();
            setTimeout(function() {
                $('#alert').fadeOut();
            }, 2000);

            return;
        }

        var formData = {
            username: $('#username').val(),
            nombre: $('#nombre').val(),
            apellido: $('#apellido').val(),
            telefono: $('#telefono').val(),
            email: $('#email').val(),
            password: password,
            confirmPassword: confirmPassword,
            paymentOption: $('input[name="payment_option"]:checked').val()
        };

        console.log('Will try to send: ');
        console.log(formData);
        $.ajax({
            url: '/create-account',
            type: 'POST',
            data: JSON.stringify(formData),
            contentType: 'application/json',
            success: function(response) {
                console.log('OK');
                console.log(response);

                // Redirigir a otra página después del registro exitoso
                // TODO: fix this...
                window.location.href = '/some-other-page';

            },
            error: function(xhr, status, error) {
                let errorMessage = 'Unknown error';
                try {
                    const responseJSON = JSON.parse(xhr.responseText);
                    errorMessage = responseJSON.message || errorMessage;
                } catch (e) {
                    console.error('Error parsing JSON:', e);
                }

                $('#alert').text('Error en el registro: ' + errorMessage).fadeIn();
                setTimeout(function() {
                    $('#alert').fadeOut();
                }, 2000);
            }
        });
    });

});

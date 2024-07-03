$(document).ready(function() {
    $('#signUp').on('submit', function(e) {
        e.preventDefault();

        console.log('I am here...');

        var password = $('#password').val();
        var confirmPassword = $('#confirm_password').val();

        if (password !== confirmPassword) {
            $('#alert').fadeIn();
            console.log('They do not match');
            setTimeout(function() {
                $('#alert').fadeOut();
            }, 2000);

            return;
        } else {
            console.log('They do match :), all good');
        }

        var formData = {
            nombre: $('#nombre').val(),
            apellido: $('#apellido').val(),
            telefono: $('#telefono').val(),
            email: $('#email').val(),
            password: $('#password').val(),
            confirm_password: $('#confirm_password').val(),
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

                // Redirigir a otra página después del registro exitoso
                window.location.href = '/some-other-page';

            },
            error: function(xhr, status, error) {
                console.error('Error:', error);
                $('#alert').text('Error en el registro: ' + error).fadeIn();
                setTimeout(function() {
                    $('#alert').fadeOut();
                }, 2000);
            }
        });
    });

});

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
            nombre: $('#nombre').val(),
            apellido: $('#apellido').val(),
            telefono: $('#telefono').val(),
            email: $('#email').val(),
            password: password,
            confirmPassword: confirmPassword,
        };

        console.log('Will try to send: ');
        console.log(formData);
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

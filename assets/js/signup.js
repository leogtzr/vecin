$(document).ready(function() {
    function isValidPhoneNumber(phone) {
        const phoneRegex = /^\+?[1-9]\d{10,10}$/;
        return phoneRegex.test(phone);
    }

    function isValidPassword(password) {
        const minLength = 8;
        const hasUpperCase = /[A-Z]/.test(password);
        const hasLowerCase = /[a-z]/.test(password);
        const hasNumbers = /\d/.test(password);
        const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(password);

        return password.length >= minLength && hasUpperCase && hasLowerCase && hasNumbers && hasSpecialChar;
    }

    $('#telefono').on('input', function() {
        if ($(this).val() && !isValidPhoneNumber($(this).val())) {
            $(this).addClass('is-invalid');
            $('#phoneHelp').text('Por favor, ingrese un número de teléfono válido.').show();
        } else {
            $(this).removeClass('is-invalid').addClass('is-valid');
            $('#phoneHelp').hide();
        }
    });

    $('#signUp').on('submit', function(e) {
        e.preventDefault();

        var password = $('#password').val();
        var confirmPassword = $('#confirm_password').val();
        var phone = $('#telefono').val();

        if (password !== confirmPassword) {
            $('#alert').fadeIn();
            setTimeout(function() {
                $('#alert').fadeOut();
            }, 2000);

            return;
        }

        var email = $('#email').val();

        if (phone && !isValidPhoneNumber(phone)) {
            $('#alert').text('Por favor, ingrese un número de teléfono válido.').fadeIn();
            setTimeout(function() {
                $('#alert').fadeOut();
            }, 6000);
            return;
        }

        // var formData = {
        //     username: $('#username').val(),
        //     nombre: $('#nombre').val(),
        //     apellido: $('#apellido').val(),
        //     telefono: $('#telefono').val(),
        //     email: email,
        //     password: password,
        //     confirmPassword: confirmPassword,
        //     paymentOption: $('input[name="payment_option"]:checked').val()
        // };

        $.ajax({
            url: '/check-email',
            type: 'POST',
            data: JSON.stringify({ email: email }),
            contentType: 'application/json',
            success: function(response) {
                if (response.exists) {
                    $('#alert').text('El email ya está registrado').fadeIn();
                    setTimeout(function() {
                        $('#alert').fadeOut();
                    }, 6000);
                } else {
                    var formData = {
                        username: $('#username').val(),
                        nombre: $('#nombre').val(),
                        apellido: $('#apellido').val(),
                        telefono: phone,
                        email: email,
                        password: password,
                        confirmPassword: confirmPassword,
                        // paymentOption: $('input[name="payment_option"]:checked').val()
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

                            window.location.href = '/confirm-account-pending';
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
                }
            },
            error: function(xhr, status, error) {
                let errorMessage = 'Unknown error';
                try {
                    const responseJSON = JSON.parse(xhr.responseText);
                    errorMessage = responseJSON.message || errorMessage;
                } catch (e) {
                    console.error('Error parsing JSON:', e);
                }

                $('#alert').text('Error en la verificación del email: ' + errorMessage).fadeIn();
                setTimeout(function() {
                    $('#alert').fadeOut();
                }, 2000);
            }
        });
    });

});

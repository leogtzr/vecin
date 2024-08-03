$(document).ready(function() {
    $('#signIn').on('submit', function(e) {
        e.preventDefault();

        var formData = {
            email: $('#email').val(),
            password: $('#password').val(),
        };

        $.ajax({
            url: '/signIn',
            type: 'POST',
            data: JSON.stringify(formData),
            contentType: 'application/json',
            success: function(response) {
                console.log('OK');
                console.log(response);

                window.location.href = '/dashboard';
            },
            error: function(xhr, status, error) {
                console.log(status);
                console.log(error);
                let errorMessage = 'Unknown error';
                const alertBox = $('#alert');
                const resendLink = $('#resend-link');

                try {
                    const responseJSON = JSON.parse(xhr.responseText);
                    errorMessage = responseJSON.message || errorMessage;
                    $('#alert-message').text(errorMessage).fadeIn();

                    // Mostrar el enlace para reenviar el correo si la cuenta no está activada
                    if (errorMessage.includes("Tu cuenta no ha sido activada")) {
                        resendLink.show();
                        resendLink.off('click').on('click', function() {
                            resendActivationEmail(formData.email);
                        });
                    } else {
                        resendLink.hide();
                    }

                    alertBox.fadeIn();
                    setTimeout(function() {
                        $('#alert').fadeOut();
                    }, 10000);
                } catch (e) {
                    console.error('Error parsing JSON:', e);
                }
            }
        });
    });

    function resendActivationEmail(email) {
        $.ajax({
            url: '/resend-activation',
            type: 'POST',
            data: JSON.stringify({ email: email }),
            contentType: 'application/json',
            success: function(response) {
                console.log(response);

                window.location.href = '/confirm-account-pending';
            },
            error: function(xhr, status, error) {
                console.error('Error:', error);
                console.log(xhr);
                console.log(status);

                let errorMessage = "Error al reenviar el correo de activación. Por favor, inténtalo de nuevo más tarde.";
                $('#alert-message').text(errorMessage).fadeIn();
            }
        });
    }
});
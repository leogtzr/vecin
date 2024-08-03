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
                try {
                    const responseJSON = JSON.parse(xhr.responseText);
                    errorMessage = responseJSON.message || errorMessage;
                    $('#alert').text(errorMessage).fadeIn();
                    setTimeout(function() {
                        $('#alert').fadeOut();
                    }, 3000);
                } catch (e) {
                    console.error('Error parsing JSON:', e);
                }

                /*
                $('#alert').text('Error en el registro: ' + errorMessage).fadeIn();
                setTimeout(function() {
                    $('#alert').fadeOut();
                }, 2000);*/
            }
        });
    });
});
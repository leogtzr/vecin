// modalHelper.js

export function showModal(message, type = 'error') {
    let emoji, title, imageUrl;

    switch(type) {
        case 'success':
            emoji = '✅';
            title = 'Completado';
            imageUrl = '/assets/images/success.gif';
            break;
        case 'warning':
            emoji = '⚠️';
            title = 'Advertencia';
            imageUrl = '/assets/images/warning.png';
            break;
        case 'error':
        default:
            emoji = '❌';
            title = 'Error 😢';
            imageUrl = '/assets/images/warn-error.gif';
            break;
    }

    $('#modalEmoji').text(emoji);
    $('#modalTitle').text(title);
    $('#errorModalBody').text(message);

    if (imageUrl) {
        $('#modalImage').attr('src', imageUrl).show();
    } else {
        $('#modalImage').hide();
    }

    $('#errorModal').modal('show');
}

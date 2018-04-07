var currentDate;
var monthNames = [ "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December" ];

$(function () {
    $('#content').width($('#content').height() / 16 * 10);
    currentDate = new Date();
    updateCurrentDate();
});

$(window).resize(function () {
    $('#content').width($('#content').height() / 16 * 10);
});

function initializeHours() {
    for (var i = 0; i < 24; i++) {
        $('#hour-view').append('<div class="hour"><span class="time">' + (i < 10 ? i + "0" : i) + ':00</span></div>');
    }

    $('#hour-view').append('<div class="hour"><span class="time">00:00</span></div>');
}

/* Agenda interaction */
$('#agenda .prev-date').on('click', function () {
    currentDate = new Date(currentDate.setTime(currentDate.getTime() - 86400000));
    updateCurrentDate();
});

$('#agenda .next-date').on('click', function () {
    currentDate = new Date(currentDate.setTime(currentDate.getTime() + 86400000));
    updateCurrentDate();
});

$('#agenda .this-date').on('click', function () {
    currentDate = new Date();
    updateCurrentDate();
})

function updateCurrentDate() {
    $('#agenda .current-date').text(currentDate.getDate() + " " + monthNames[currentDate.getMonth()] + " " + currentDate.getFullYear());
}
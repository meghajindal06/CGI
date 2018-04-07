var resourceIndex = 0, resourceSpeed = 4, spawnBaseSpeed = 3000, nodes = 2;
var resourceTypes = ['energy', 'people', 'goods'];

$(function () {
    $('#network-value .value').text((1 * (Math.pow(nodes, 1.69) * 3)).toFixed(3));
    DrawLines();
    StartAnimation();
});

/* Simulation controls */
$('#controls .add-node').on('click', function () {
    nodes++;
    console.log(nodes);

    $('#diagram').append('<img id = "node' + nodes + '" class= "element node" src = "img/node.png" />');
    var angle = (2 * Math.PI) / nodes;
    var radius = 350;
    for (var i = 1; i <= nodes; i++) {
        var xPos = radius * Math.sin(angle * i);
        var yPos = radius * Math.cos(angle * i);
        $('#node' + i).css('left', (xPos + $('#diagram').width() / 2) + 'px').css('top', (yPos + $('#diagram').height() / 2) + 'px');
    }

    $('#network-value .value').text((1 * (Math.pow(nodes, 1.69) * 3)).toFixed(3));

    ClearLines();
    DrawLines();
    ForceClearResources();
});

/* Canvas manipulation */
function DrawLines() {
    var canvas = document.getElementById("background"),
        ctx = canvas.getContext("2d"),
        offset = 38;

    $('#diagram .element.node').each(function () {
        var myId = $(this).attr('id'),
            myPosition = $(this).position();

        $('#diagram .element.node').each(function () {
            if ($(this).attr('id') != myId) {
                ctx.beginPath();
                ctx.moveTo(myPosition.left + offset, myPosition.top + offset);
                ctx.lineTo($(this).position().left + offset, $(this).position().top + offset);
                ctx.strokeStyle = '#C1C9CB';
                ctx.stroke();
            }
        });
    });
}

function ClearLines() {
    var canvas = document.getElementById("background"),
        ctx = canvas.getContext("2d");

    ctx.clearRect(0, 0, canvas.width, canvas.height);
}

/* Animations */
function StartAnimation() {
    var start = resourceIndex;
    setTimeout(function () {
        ClearResource(start, 1);
    }, resourceSpeed * 1000 + 250);

    var origin = getRandomInt(1, nodes),
        target = getRandomInt(1, nodes);
    while (origin == target) {
        target = getRandomInt(1, nodes);
    }

    SpawnResource('node' + origin, 'node' + target, getRandomInt(0, resourceTypes.length));

    setTimeout(function () {
        StartAnimation();
    }, spawnBaseSpeed / (nodes * 2));
}

function SpawnResource(startElement, targetElement, type) {
    $('#diagram').append("<img id='resource-" + resourceIndex + "' class='resource element' src='img/" + resourceTypes[type] + ".png' />");
    var offset = 18;

    var res = $('#resource-' + resourceIndex);
    res.css('top', $('#' + startElement).position().top + offset);
    res.css('left', $('#' + startElement).position().left + offset);
    res.css('transition', resourceSpeed + 's');
    res.css('top', $('#' + targetElement).position().top + offset);
    res.css('left', $('#' + targetElement).position().left + offset);

    resourceIndex++;
}

function ClearResource(resourceID) {
    $('#diagram #resource-' + resourceID).remove();
}

function ForceClearResources() {
    $('#diagram .resource.element').remove();
}

/* Utility functions */
function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}
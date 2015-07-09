var main = function() {
    $('h1').click(function() {
        toggle($(this), 'active');

        $('h1').not(this).each(function() {
            $(this).removeClass('active');
            $(this).siblings().each(function() {
                $(this).css('color', $(this).siblings().css('color'));
            });
        });
    });
}

var toggle = function(element, cssClass) {
    if(element.hasClass(cssClass))
        element.removeClass(cssClass);
    else
        element.addClass(cssClass);

    element.siblings().each(function() {
        $(this).css('color', $(this).siblings().css('color'));
    });
}

var colorLinks = function() {
    // Colors each link like the header

    $('a').each(function() {
        $(this).css('color', $(this).siblings().css('color'));
    });
}

var textShadowHeaders = function() {
    $('h1').each(function() {
        $(this).css('text-shadow', '0 0 4px ' + $(this).css('color'));
    });
}

var writeLinks = function() {
    var links = getLinks();

    $('a').each(function(index) {
        $(this).attr('href', links[index]);
    });
}

var writeLinkNames = function() {
    var linkNames = getLinkNames();

    $('a').each(function(index) {
        $(this).text(linkNames[index]);
    });
}

var writeHeaders = function() {
    var headers = getHeaders();

    $('h1').each(function(index) {
        $(this).text(headers[index]);
    });
}

var writeContent = function() {
    writeHeaders();
    writeLinkNames();
    writeLinks();
}

var init = function() {
    $(document).ready(writeContent);
    $(document).ready(main);
    $(document).ready(colorLinks);
    $(document).ready(textShadowHeaders);
}

init();

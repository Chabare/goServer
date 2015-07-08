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

    $('body').css('background-image', 'url(\'http://rpi.eowqofhvla66ahgq.myfritz.net/gruende/img/background.png\')');
}

var textShadowHeaders = function() {
    $('h1').each(function() {
        $(this).css('text-shadow', '0 0 4px ' + $(this).css('color'));
    });
}

var getLinks = function() {
    return [
        'http://bitbucket.com/',
        'http://www.drive.google.com/',
        'http://github.com',
        'http://www.animehaven.org/',
        'http://www.crunchyroll.com/home/queue',
        'http://myanimelist.net/animelist/chabare',
        'http://www.congstar.com/',
        'http://www.meine-ksk.de/',
        'http://www.paypal.de/',
        'http://www.theverge.com/',
        'http://www.theverge.com/',
        'http://www.wired.com/',
        'http://www.sight-board.de/tu-darmstadt/',
        'https://moodle.informatik.tu-darmstadt.de/',
        'https://www.tucan.tu-darmstadt.de/',
        'http://www.amazon.de/',
        'http://www.humblebundle.com/',
        'http://store.steampowered.com/',
        'http://www.4chan.org/b/',
        'http://www.reddit.com/',
        'https://web.whatsapp.com/',
        'http://www.netflix.com/',
        'http://www.twitch.tv/directory/following',
        'http://www.youtube.com/feed/subscriptions',
        'http://www.codeacademy.com',
        'http://www.forecast.io',
        'http://www.imgur.com'
    ];
}

var writeLinks = function() {
    var links = getLinks();

    $('a').each(function(index) {
        $(this).attr('href', links[index]);
    });
}

var getLinkNames = function() {
    return [
        'BitBucket',
        'Drive',
        'Github',
        'Animehaven',
        'Crunchyroll',
        'MyAnimeList',
        'Congstar',
        'KSK',
        'PayPal',
        'ExtremeTech',
        'TheVerge',
        'Wired',
        'Campus-Navi',
        'Moodle',
        'TuCan',
        'Amazon',
        'HumbleBundle',
        'Steam',
        '4Chan',
        'Reddit',
        'WhatsApp',
        'Netflix',
        'Twitch',
        'YouTube',
        'CodeAcademy',
        'Forecast',
        'Imgur'
    ];
}

var writeLinkNames = function() {
    var linkNames = getLinkNames();

    $('a').each(function(index) {
        $(this).text(linkNames[index]);
    });
}

var getHeaders = function() {
    return [
        'Clouds',
        'Anime',
        'Money',
        'News',
        'University',
        'Shops',
        'Social',
        'Streaming',
        'Stuff'
    ];
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

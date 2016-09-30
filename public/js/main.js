/**
 * Created by johnla on 2016-09-26.
 */

$(document).ready(function () {
    $("#main-menu-button").on('click', function () {
        $("#menu > ul").slideToggle('medium', function () {
            if($(this).is(':visible'))
                $(this).css('display', 'flex');
        });
    })
});
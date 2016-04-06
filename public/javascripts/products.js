'use strict';

$(function(){


    $('.products__gallery--thumbs').flexslider({
        animation: "slide",
        controlNav: false,
        animationLoop: false,
        slideshow: false,
        itemWidth: 80,
        itemMargin: 16,
        asNavFor: '.products__gallery--top'
    });

    $('.products__gallery--top').flexslider({
        animation: "slide",
        controlNav: false,
        directionNav: false,
        animationLoop: false,
        slideshow: false,
        sync: ".products__gallery--thumbs"
    });



    $('.products__featured--slider').flexslider({
        animation: "slide",
        animationLoop: false,
        controlNav: false,
        itemWidth: 238,
        itemMargin: 16
    });
})

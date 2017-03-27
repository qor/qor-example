'use strict';

$(function(){
  let qty_input = $('.products__meta--qty input').get(0);
  $('.products__meta--qty a.reduce').click(function(event) {
    event.preventDefault();
    qty_input.stepDown();
  });
  $('.products__meta--qty a.add').click(function(event) {
    event.preventDefault();
    qty_input.stepUp();
  });

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

  let productsFeaturedSliderH = $('.products__featured--slider').width();
  let isMobile = window.matchMedia("only screen and (max-width: 760px)").matches;
  let columnNuber = isMobile ? 2 : 4;

  $('.products__featured--slider').flexslider({
    animation: "slide",
    animationLoop: false,
    controlNav: false,
    itemWidth: (productsFeaturedSliderH - 16 * 3) / columnNuber,
    itemMargin: 16
  });
})

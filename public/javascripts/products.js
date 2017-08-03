"use strict";

$(function() {
  let $qty = $(".products__meta--qty");
  $qty.find("a.reduce").on("click", function(event) {
    event.preventDefault();
    $qty.find("input").get(0).stepDown();
  });
  $qty.find("a.add").on("click", function(event) {
    event.preventDefault();
    $qty.find("input").get(0).stepUp();
  });

  $(".products__details").on("submit", function(event) {
    event.preventDefault();
    $.ajax({
      type: "POST",
      url: "/cart/",
      error: function(xhr) {
        alert(xhr.status + ": " + xhr.statusText);
      },
      success: function(response) {
        alert(
          response.status +
            ": " +
            response.message +
            " (" +
            response.itemID +
            ")"
        );
      },
      data: $(event.target).serialize()
    });
  });

  $(".products__gallery--thumbs").length &&
    $(".products__gallery--thumbs").flexslider({
      animation: "slide",
      controlNav: false,
      animationLoop: false,
      slideshow: false,
      itemWidth: 80,
      itemMargin: 16,
      asNavFor: ".products__gallery--top"
    });

  $(".products__gallery--top").length &&
    $(".products__gallery--top").flexslider({
      animation: "slide",
      controlNav: false,
      directionNav: false,
      animationLoop: false,
      slideshow: false,
      sync: ".products__gallery--thumbs"
    });

  let productsFeaturedSliderH = $(".products__featured--slider").width(),
    isMobile = window.matchMedia("only screen and (max-width: 760px)").matches,
    columnNuber = isMobile ? 2 : 4;

  $(".products__featured--slider").length &&
    $(".products__featured--slider").flexslider({
      animation: "slide",
      animationLoop: false,
      controlNav: false,
      itemWidth: 200,
      itemMargin: 16
    });
});

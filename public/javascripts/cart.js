'use strict';

$(function() {
  $(".cart__list--remove").click(function(event) {
    event.preventDefault();

    $.ajax({
      type: "DELETE",
      url: event.target.href,
      success: function() {
        location.reload();
      }
    });
  });
});

'use strict';

$(function() {
  $('.products__details').submit(function(event) {
    event.preventDefault();
    $.ajax({
      type: "POST",
      url: "/cart/",
      error: function(xhr) {
        alert(xhr.status + ": " + xhr.statusText);
      },
      success: function(response) {
        alert(response.status + ": " + response.message + " (" + response.itemID + ")");
      },
      data: $(event.target).serialize()
    });
});

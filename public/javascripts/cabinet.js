'use strict'

$(function(){

  $("#billing-address").submit(function(event) {
    event.preventDefault();
    $.ajax({
      type: "POST",
      url: "/cabinet/billing_address",
      error: function(xhr) {
        alert(xhr.status + ": " + xhr.statusText);
      },
      success: function(response) {
        alert(response.status + ": " + response.message + " (" + response.itemID + ")");
      },
      data: $(event.target).serialize()
    });
  });

  $("#shipping-address").submit(function(event) {
    event.preventDefault();
    $.ajax({
      type: "POST",
      url: "/cabinet/shipping_address",
      error: function(xhr) {
        alert(xhr.status + ": " + xhr.statusText);
      },
      success: function(response) {
        alert(response.status + ": " + response.message + " (" + response.itemID + ")");
      },
      data: $(event.target).serialize()
    });
  });
})

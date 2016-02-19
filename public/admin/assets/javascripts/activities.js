$(document).ready(function() {
  $(".qor-field [type=submit]").parents("form").submit(function(event) {
    $.post($(this).attr("action"), $(this).serialize(), function(data) {
      console.log(data)
    })
    return false;
  });
})

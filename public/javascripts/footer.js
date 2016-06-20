'use strict';

$(function(){
    $('.footer__change-language select').on('change', function (){
        var url = $(this).val();
          if (url) {
              window.location = url;
          }
          return false;
    });

})

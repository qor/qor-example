'use strict';

function updateCart(data) {
    $.ajax({
        type: 'PUT',
        data: data,
        dataType: 'json',
        url: '/cart',
        beforeSend: function() {
            $('.cart__qty--btn').attr('disabled', true);
        },
        success: function() {
            window.location.reload();
        },
        error: function(xhr) {
            window.alert(xhr.status + ': ' + xhr.statusText);
            window.location.reload();
        }
    });
}

$(function() {
    $('.cart__list--remove').on('click', function(e) {
        e.preventDefault();
        let $num = $(this)
            .closest('tr')
            .find('.cart__qty--num');

        updateCart(
            {
                quantity: 0,
                size_variation_id: $num.data('size-variation-id')
            }
        );
    });

    $('.cart__qty--minus').on('click', function(e) {
        let $num = $(this)
                .attr('disabled', true)
                .closest('.cart__qty')
                .find('.cart__qty--num'),
            colorVariationID = $num.data('size-variation-id'),
            currentVal = parseInt($num.val()),
            qty = currentVal - 1;

        if (qty < 1) {
            qty = 0;
        }

        $num.val(qty);
        updateCart(
            {
                quantity: qty,
                size_variation_id: colorVariationID
            }
        );

        return false;
    });

    $('.cart__qty--plus').on('click', function(e) {
        let $num = $(this)
                .attr('disabled', true)
                .closest('.cart__qty')
                .find('.cart__qty--num'),
            colorVariationID = $num.data('size-variation-id'),
            currentVal = parseInt($num.val()),
            qty = currentVal + 1;

        $num.val(qty);

        updateCart({
            quantity: qty,
            size_variation_id: colorVariationID
        });

        return false;
    });

    $('.cart__qty--num').on('blur', function(e) {
        let $num = $(this).attr('disabled', true),
            colorVariationID = $num.data('size-variation-id'),
            qty = $num.val();

        if (!/^[0-9]*$/.test(qty)) {
            window.alert('please enter the right number!');
            return;
        }

        updateCart({
            quantity: parseInt(qty),
            size_variation_id: colorVariationID
        });

        return false;
    });
});

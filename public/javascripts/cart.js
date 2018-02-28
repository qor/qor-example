'use strict';

function updateCart(data, type) {
    $.ajax({
        type: type ? 'DELETE' : 'PUT',
        data: data,
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
        updateCart(
            {
                qty: 0,
                product_id: '',
                color_variation_id: ''
            },
            true
        );
    });

    $('.cart__qty--minus').on('click', function(e) {
        let $num = $(this)
                .attr('disabled', true)
                .closest('.cart__qty')
                .find('.cart__qty--num'),
            currentVal = parseInt($num.val()),
            rightVal = currentVal - 1;

        if (rightVal < 1) {
            updateCart(
                {
                    qty: 0,
                    product_id: '',
                    color_variation_id: ''
                },
                true
            );

            return false;
        }

        $num.val(rightVal);

        updateCart({
            qty: rightVal,
            product_id: '',
            color_variation_id: ''
        });

        return false;
    });

    $('.cart__qty--plus').on('click', function(e) {
        let $num = $(this)
                .attr('disabled', true)
                .closest('.cart__qty')
                .find('.cart__qty--num'),
            currentVal = parseInt($num.val()),
            rightVal = currentVal + 1;

        $num.val(rightVal);

        updateCart({
            qty: rightVal,
            product_id: '',
            color_variation_id: ''
        });

        return false;
    });

    $('.cart__qty--num').on('blur', function(e) {
        let $num = $(this).attr('disabled', true),
            currentVal = $num.val();

        if (!/^[0-9]*$/.test(currentVal)) {
            window.alert('please enter the right number!');
            return;
        }

        updateCart({
            qty: parseInt(currentVal),
            product_id: '',
            color_variation_id: ''
        });

        return false;
    });
});

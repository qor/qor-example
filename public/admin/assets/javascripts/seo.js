!(function($, w) {
    "use script";

    var CLASS_SEO = ".qor-seo";
    var CLASS_SUBMIT = ".qor-seo-submit";
    var CLASS_FIELD = ".qor-seo-field";
    var CLASS_TITLE_NAME = ".qor-seo-title-field";
    var CLASS_DESCRIPTION_NAME = ".qor-seo-description-field";
    var CLASS_TAGS_NAME = ".qor-seo-tags-field";
    var CLASS_ADD_TAGS_NAME = ".qor-seo-tag";
    var CLASS_TAGS_INPUT_NAME = ".qor-seo-input-field";

    function QorSeo($element){
        this.$element = $element;
        this.focusedInputID = "";
    }
    QorSeo.prototype={
        init: function(){
            this.$wrap = this.$element.parents(CLASS_SEO);
            this.$addTgas = this.$wrap.find(CLASS_ADD_TAGS_NAME);
            this.$tagInputs = this.$wrap.find(CLASS_TAGS_INPUT_NAME);
            this.bind();
        },
        bind: function () {
            this.$element.on('click', $.proxy(this.submitSeo, this));
            this.$tagInputs.on('click keyup', $.proxy(this.tagInputsFocus, this));
            this.$tagInputs.on('blur', $.proxy(this.tagInputsBlur, this));
            this.$addTgas.on('click', $.proxy(this.addTags, this));
        },
        saveSelection: function(containerEl) {
            var range = window.getSelection().getRangeAt(0);
            var preSelectionRange = range.cloneRange();
            preSelectionRange.selectNodeContents(containerEl);
            preSelectionRange.setEnd(range.startContainer, range.startOffset);
            var start = preSelectionRange.toString().length;

            return {
                start: start,
                end: start + range.toString().length
            };
        },
        tagInputsFocus: function(){
            this.$addTgas.addClass('focus');
            var $focusedInput = $(document.activeElement);

            this.focusedInputID = $focusedInput.prop("id");
            this.focusedInputStart = $focusedInput[0].selectionStart;
            this.focusedInputEnd = $focusedInput[0].selectionEnd;
            this.focusedInputVal = $focusedInput.val();
        },
        tagInputsBlur: function(){
            this.$addTgas.removeClass('focus');
            this.$focusedInputID = false;
        },
        addTags: function(event){
            if (!this.focusedInputID){
                return;
            }

            var newVal = "";
            var startString = this.focusedInputVal.substring(0,this.focusedInputStart);
            var endString = this.focusedInputVal.substring(this.focusedInputEnd,this.focusedInputVal.length);
            var tagVal = "{{"+$(event.currentTarget).data("tagValue")+"}}";

            newVal = startString + tagVal + endString;
            $("#"+this.focusedInputID).val(newVal).focus();
        },
        submitSeo: function(){
            var fieldName = this.$wrap.find(CLASS_FIELD).prop("name");
            var titleValue = this.$wrap.find(CLASS_TITLE_NAME).val();
            var tagsValue = this.$wrap.find(CLASS_TAGS_NAME).val();
            var descriptionValue = this.$wrap.find(CLASS_DESCRIPTION_NAME).val();
            var data = {};

            data[fieldName] = JSON.stringify({
                "Title": titleValue,
                "Description": descriptionValue,
                "Tags": tagsValue
            });

            var url = this.$wrap.parents(".qor-form").attr("action");
            $.ajax({
                type: "POST",
                url: url,
                data: data,
                success: function () {
                    $('.qor-alert--success').show().addClass('');
                    setTimeout(function () {
                        $('.qor-alert--success').hide();
                      }, 5000);
                },
                error: function (data) {
                    $('.qor-alert--error').show();
                }
            });
            return false;
        }
    }

    $(function(){
        var submits = $(CLASS_SUBMIT);
        submits.each(function(){
            var qorSeo = new QorSeo($(this));
            qorSeo.init();
        });

        $(document).on('click.qor.fixedAlert', '[data-dismiss="fixed-alert"]', function () {
            $(this).closest('.qor-alert').hide();
        });
    })

})(jQuery, window);

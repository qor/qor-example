!(function() {
  (function($, Export) {
    "use script";

    var StringUtils = {
      Parameterize : function(str, sep) {
        var parameterizedStr = str.replace(/[^a-z0-9\-_]+/gi, sep)

        if (sep != "") {
          // No more than one of the separator in a row.
          parameterizedStr = parameterizedStr.replace(new RegExp(sep + '{2, }', 'g'), sep)
          // Remove leading/trailing separator.
          parameterizedStr = parameterizedStr.replace(new RegExp('^' + sep + '|' + sep + '$', 'g'), '')
        }

        return parameterizedStr.toLowerCase()
      }
    };

    var testParameterize = function() {
      var cases = [
        {in: "This is an blog title", want:"this-is-an-blog-title"},
        {in: " This is an blog title with spaces  ", want:"this-is-an-blog-title-with-spaces"},
        {in: "Donald E. Knuth", want:"donald-e-knuth"},
        {in: "Two  More   Spaces", want:"two-more-spaces"},
        {in: "这是一个标题", want:""},
      ]

      for (var i = 0; i < cases.length; i++) {
        var c = cases[i];
        var got = StringUtils.Parameterize(c.in, "-");
        if (got != c.want) {
          throw new Error('got StringUtils.Parameterize("' + c.in + '") = "' + got + '", want "' + c.want + '"');
        }
      }
    };

    var Slug = {
      Init: function() {
        this.bindEvents();
      },

      bindEvents: function() {
        $(document).on('slug.change', '[data-slug="true"]', function(evt) {
          var $container = $(evt.target);
          var $source = $container.find('[data-slug-role="source"]'),
            $slug = $container.find('[data-slug-role="slug"]'),
            $sync = $container.find('[data-slug-role="sync"]');

          if ($sync.is(':checked')) {
            $slug.val(StringUtils.Parameterize($source.val(), '-'))
          }
        });

        $(document).on('keyup change', '[data-slug-role="source"]', function(evt) {
          $(evt.target).parents('[data-slug="true"]').trigger('slug.change');
        });

        $(document).on('change', '[data-slug-role="sync"]', function(evt) {
          $(evt.target).parents('[data-slug="true"]').trigger('slug.change');
        });
      }
    }

    testParameterize();
    Export.QorSlug = Slug;

  })(jQuery, window);
}).call(this);

$(function() {
  QorSlug.Init();
})

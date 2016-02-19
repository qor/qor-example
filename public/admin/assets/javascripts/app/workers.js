(function (factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as anonymous module.
    define(['jquery'], factory);
  } else if (typeof exports === 'object') {
    // Node / CommonJS
    factory(require('jquery'));
  } else {
    // Browser globals.
    factory(jQuery);
  }
})(function ($) {

  'use strict';

  var NAMESPACE = 'qor.worker';
  var EVENT_ENABLE = 'enable.' + NAMESPACE;
  var EVENT_DISABLE = 'disable.' + NAMESPACE;
  var EVENT_CLICK = 'click.' + NAMESPACE;
  var CLASS_NEW_WORKER = '.qor-worker--new';
  var CLASS_WORKER_ERRORS = '.qor-worker--show-errors';

  function QorWorker(element, options) {
    this.$element = $(element);
    this.options = $.extend({}, QorWorker.DEFAULTS, $.isPlainObject(options) && options);
    this.init();
  }

  QorWorker.prototype = {
    constructor: QorWorker,

    init: function () {
      var $this = this.$element;
      this.bind();
    },

    bind: function () {
      this.$element.on(EVENT_CLICK, $.proxy(this.click, this));
    },

    click: function (e) {
      var $target = $(e.target);
      e.stopPropagation();

      if ($target.is(CLASS_WORKER_ERRORS)){
        var $workerErrorModal = $(QorWorker.POPOVERTEMPLATE).appendTo('body');
        var url = $('tr.is-selected .qor-button--edit').attr('href');
        $workerErrorModal.qorModal('show');

        $.ajax({
          url: url
        }).done(function (html) {
          var $content = $(html).find('.qor-form-container');
          var $errorTable = $content.find('.workers-error-output');
          if ($errorTable){
            $errorTable.appendTo($workerErrorModal.find('#qor-worker-errors'));
          }
        });
      }

      if ($target.is(CLASS_NEW_WORKER)){
        var $targetParent = $target.parent();

        $targetParent.addClass('current');
        $('.qor-worker-form-list').not('current').find('form').addClass('hidden');
        $target.next('form').toggleClass('hidden');
      }
    }
  };

  QorWorker.DEFAULTS = {};
  QorWorker.POPOVERTEMPLATE = (
     '<div class="qor-modal fade qor-modal--worker-errors" tabindex="-1" role="dialog" aria-hidden="true">' +
      '<div class="mdl-card mdl-shadow--2dp" role="document">' +
        '<div class="mdl-card__title">' +
          '<h2 class="mdl-card__title-text">Process Errors</h2>' +
        '</div>' +
        '<div class="mdl-card__supporting-text" id="qor-worker-errors"></div>' +
        '<div class="mdl-card__actions">' +
          '<a class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" data-dismiss="modal">close</a>' +
        '</div>' +
      '</div>' +
    '</div>'
  );

  QorWorker.plugin = function (options) {
    return this.each(function () {
      var $this = $(this);
      var data = $this.data(NAMESPACE);
      var fn;

      if (!data) {

        if (/destroy/.test(options)) {
          return;
        }

        $this.data(NAMESPACE, (data = new QorWorker(this, options)));
      }

      if (typeof options === 'string' && $.isFunction(fn = data[options])) {
        fn.apply(data);
      }
    });
  };

  $.fn.qorSliderAfterShow = function (url) {
    QorWorker.getWorkerProgressIntervId = window.setInterval(QorWorker.getWorkerProgress, 1000, url);
  };

  QorWorker.isScrollToBottom = function (element) {
    return element.clientHeight + element.scrollTop === element.scrollHeight;
  };

  QorWorker.getWorkerProgress = function (url) {

    var progressURL = url;
    var $logContainer = $('.workers-log-output');
    var $progressValue = $('.qor-worker--progress-value');
    var $progressStatusStatus = $('.qor-worker--progress-status');
    var workerProgress = document.querySelector('#qor-worker--progress');

    if ($('.qor-worker--progress').data('worker-progress') == 100){
      window.clearInterval(QorWorker.getWorkerProgressIntervId);
      document.querySelector('#qor-worker--progress').MaterialProgress.setProgress(100);
      return;
    }

    $.ajax({
      url: progressURL
    }).done(function (html) {
      var $content = $(html).find('.qor-form-container');
      var currentStatus = $content.find('.qor-worker--progress').data('worker-progress');
      var progressStatusStatus = $content.find('.qor-worker--progress').data('worker-status');

      if (!currentStatus){
        window.clearInterval(QorWorker.getWorkerProgressIntervId);
        return;
      }

      $progressValue.html(currentStatus);
      $progressStatusStatus.html(progressStatusStatus);

      // set status progress
      if (workerProgress && workerProgress.MaterialProgress){
        workerProgress.MaterialProgress.setProgress(currentStatus);
      }

      if (currentStatus >= 100){
        window.clearInterval(QorWorker.getWorkerProgressIntervId);
        return;
      }
      if (currentStatus < 100){
        // update process log
        var oldLog = $.trim($logContainer.html());
        var newLog = $.trim($content.find('.workers-log-output').html());
        var newLogHtml;

        if (newLog != oldLog){
          newLogHtml = newLog.replace(oldLog, '');

          if (QorWorker.isScrollToBottom($logContainer[0])){
            $logContainer.append(newLogHtml).scrollTop($logContainer[0].scrollHeight);
          } else {
            $logContainer.append(newLogHtml);
          }

        }
      }
    });
  };

  $(function () {
    var selector = '[data-toggle="qor.workers"]';

    $(document).
      on(EVENT_DISABLE, function (e) {
        QorWorker.plugin.call($(selector, e.target), 'destroy');
      }).
      on(EVENT_ENABLE, function (e) {
        QorWorker.plugin.call($(selector, e.target));
      }).
      triggerHandler(EVENT_ENABLE);
  });

  return QorWorker;

});

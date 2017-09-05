define(function(require, exports, module) {
  var $ = require('jquery'),
    bootstrap = require('bootstrap'),
    global = require('global');

  $(function() {
    $.ajax({
      type: "GET",
      url: "/ajaxdetail",
      success: function(obj) {
        console.log(obj);
        $(".product-price").html(obj.Price);
        $(".product-title").html(obj.Title);
      },
      error: function(err) {
        alert("something is wrong in webapp");
      },
    });
  });
});

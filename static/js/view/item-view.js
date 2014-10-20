/*global Backbone, jQuery, _, ENTER_KEY, ESC_KEY */
var app = app || {};

(function ($) {
	'use strict';

	app.ItemView = Backbone.View.extend({
		tagName:  'li',

		template: _.template($('#item-template').html()),
		//monsterTemplate: _.template($('#monster-template').html()),

		render: function () {
			this.$el.html(this.template(this.model.toJSON()));
			return this;
		}
	});
})(jQuery);
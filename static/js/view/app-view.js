/*global Backbone, jQuery, _, ENTER_KEY */
var app = app || {};

(function ($) {
	'use strict';

	// The Application
	// ---------------

	// Our overall **AppView** is the top-level piece of UI.
	app.AppView = Backbone.View.extend({

		// Instead of generating a new element, bind to the existing skeleton of
		// the App already present in the HTML.
		el: '.codex',

		monsterTemplate: _.template($('#monster-template').html()),

		// Delegated events for creating new items, and clearing completed ones.
		events: {
			'click .monster-name': 'updateMonster'
		},

		initialize: function () {
			this.$list = $('.monster-list');
			this.$monster = $('.monster');

			this.listenTo(app.monsters, 'add', this.addOne);
			this.listenTo(app.monsters, 'reset', this.addAll);
			this.listenTo(app.monsters, 'all', this.render);

			app.monsters.fetch({reset: true});
		},

		render: function () {
			
		},

		addOne: function (monster) {
			var view = new app.ItemView({ model: monster });
			this.$list.append(view.render().el);
		},

		addAll: function () {
			var monsters = app.monsters;

			this.$list.html('');
			monsters.each(this.addOne, this);
			this.$monster.html(
				this.monsterTemplate(
					monsters.at(Math.floor(Math.random() * (monsters.length - 0)) + 0).toJSON()
				)
			);
		},

		updateMonster: function(e) {
			e.preventDefault();
			
			this.$monster.html(
				this.monsterTemplate(
					app.monsters.get(
						$(e.currentTarget)
						.data("id")
					).toJSON()
				)
			);
		}
	});
})(jQuery);
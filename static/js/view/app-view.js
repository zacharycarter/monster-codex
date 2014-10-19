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

		// Delegated events for creating new items, and clearing completed ones.
		events: {
			/*'keypress #new-todo': 'createOnEnter',
			'click #clear-completed': 'clearCompleted',
			'click #toggle-all': 'toggleAllComplete'*/
		},

		// At initialization we bind to the relevant events on the `Todos`
		// collection, when items are added or changed. Kick things off by
		// loading any preexisting todos that might be saved in *localStorage*.
		initialize: function () {
			// this.allCheckbox = this.$('#toggle-all')[0];
			// this.$input = this.$('#new-todo');
			// this.$footer = this.$('#footer');
			// this.$main = this.$('#main');
			this.$list = $('.monster-list');

			this.listenTo(app.monsters, 'add', this.addOne);
			this.listenTo(app.monsters, 'reset', this.addAll);
			this.listenTo(app.monsters, 'all', this.render);

			// Suppresses 'add' events with {reset: true} and prevents the app view
			// from being re-rendered for every model. Only renders when the 'reset'
			// event is triggered at the end of the fetch.
			app.monsters.fetch({reset: true});
		},

		// Re-rendering the App just means refreshing the statistics -- the rest
		// of the app doesn't change.
		render: function () {
		},

		// Add a single todo item to the list by creating a view for it, and
		// appending its element to the `<ul>`.
		addOne: function (monster) {
			var view = new app.MonsterView({ model: monster });
			this.$list.append(view.render().el);
		},

		// Add all items in the **Todos** collection at once.
		addAll: function () {
			this.$list.html('');
			app.monsters.each(this.addOne, this);
		}
	});
})(jQuery);
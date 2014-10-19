/*global Backbone */
var app = app || {};

(function () {
	'use strict';

	// Monster Collection
	// ---------------

	var Monsters = Backbone.Collection.extend({
		// Reference to this collection's model.
		model: app.Monster,

		url: "/api/monsters"
	});

	// Create our global collection of **Todos**.
	app.monsters = new Monsters();
})();
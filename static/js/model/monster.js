/*global Backbone */
var app = app || {};

(function () {
	'use strict';

	// Monster Model
	// ----------

	// Our basic **Monster** model has `name`, `keywords`, 
	// and `description` attributes.
	app.Monster = Backbone.Model.extend({
		// Default attributes for the todo
		// and ensure that each todo created has `title` and `completed` keys.
		defaults: {
			Name: '',
			Keywords: [],
			Description: ''
		}
	});
})();
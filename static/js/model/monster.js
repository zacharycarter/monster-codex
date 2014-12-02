/*global Backbone */
var app = app || {};

(function () {
	'use strict';

	// Monster Model
	// ----------

	// Our basic **Monster** model has `name`, `keywords`, 
	// and `description` attributes.
	app.Monster = Backbone.Model.extend({
		idAttribute: "Id",

		defaults: {
			Name: '',
			Talents: [],
			Keywords: [],
			Description: ''
		}
	});
})();
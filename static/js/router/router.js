/*global Backbone */
var app = app || {};

(function () {
	'use strict';

	// Monster Router
	// ----------
	var MonsterRouter = Backbone.Router.extend({
		routes: {
		}
	});

	app.MonsterRouter = new MonsterRouter();
	Backbone.history.start();
})();
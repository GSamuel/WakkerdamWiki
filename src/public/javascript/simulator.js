var gameSettings = {
	updateTotal : function(settings) {
		var total = 0;
		for (var i = 0; i < settings.cards.length; i++) {
			total += settings.cards[i].count;
		}
		settings.total = total;

		while(settings.player_names.length < settings.total) {
			settings.player_names.push("Player " + (settings.player_names.length + 1));
		}

		return settings;
	},
	saveSettings : function(settings) {
		settings = gameSettings.updateTotal(settings);
		var data = JSON.stringify(settings);
		localStorage.setItem("settings", data);
		view.onSettingsChanged(settings);
	},
	retreiveSettings : function() {
		var data = localStorage.getItem("settings");
		return JSON.parse(data);
	},
	updateAvailableCards : function(availableCards) {
		var settings = gameSettings.retreiveSettings();
		if(settings == null) {
			settings = {player_names : [], cards: [], total:0};
		}

		var cards = [];

		for (var i = 0; i < availableCards.cards.length; i++) {
			var card = availableCards.cards[i];
			card.count = 0;

			for (var j = 0; j < settings.cards.length; j++) {
				if(settings.cards[j].card_name == card.card_name) {
					var count = settings.cards[j].count;
					card.count = count;
				}
			}
			cards.push(card);
		}
		settings.cards = cards;

		gameSettings.saveSettings(settings);
	},
}

var gameState = {
	saveState : function(state) {
		var data = JSON.stringify(state);
		localStorage.setItem("gameState", data);
		view.onGameStateChanged(state);
	},
	retreiveState : function() {
		var data = localStorage.getItem("gameState");
		return JSON.parse(data);
	},
	updateCard: function(id, name){
		if(name == "unknown") {
			return ;
		}

		var state = gameState.retreiveState();
		var playerIndex = -1;

		for (var i = 0; i < state.players.length; i++) {
			if(state.players[i].id == id) {
				playerIndex = i;
			}
		}

		if(state.players[playerIndex].card.name != "unknown") {
			console.log("already has card");
			return;
		}

		for (var i = 0; i < state.unassigned.length; i++) {
			if(state.unassigned[i].name == name) {
				state.players[playerIndex].card = state.unassigned[i];
				state.unassigned.splice(i, 1);
				break;
			}
		}
		gameState.saveState(state);
	}
}

var removeChildren = function(element) {
	$(element).children().each (function(index, el) {
		if(!$(el).hasClass("prototype") && !$(el).hasClass("pinned")) {
			$(el).remove();
		}
	})
}

var view = {
	updateNames : function(settings) {
		removeChildren('#updateNamesForm');
		var names = settings.player_names;
		for (var i = 0; i < settings.total; i++) {
			var id = i +1;
			var prototype = $('#updateNamesForm .prototype')
			var clone = $(prototype).clone().insertBefore( prototype ).toggleClass("prototype");
			$(clone).find("label").html("ID: " + (id)).prop('for', "playername-"+id);
			var input = $(clone).find("input").prop("id", "playername-"+id).prop('placeholder', names[i]).prop("playerId", id).val(names[i]);
			$(input).change (function(){
				var localSettings = gameSettings.retreiveSettings();
				var index = parseInt($(this).prop("playerId")) -1;
				var value = $(this).val();
				if(localSettings.player_names[index] != value) {
					localSettings.player_names[index] = value;
					gameSettings.saveSettings(localSettings);
				}
			})
		}
	},
	onSettingsChanged : function(settings) {
		removeChildren('#newGameForm');
		for (var i = 0; i < settings.cards.length; i++) {
			var card = settings.cards[i];
			var prototype = $('#newGameForm .prototype');
			var clone = $(prototype).clone().insertBefore( prototype ).toggleClass("prototype");
			$(clone).find("label").html(card.card_name).prop("for", card.card_name);
			var select = $(clone).find("select").prop("id", card.card_name);
			for (var j = 0; j <= card.max_amount; j++) {
				$(select).append("<option>" + j + "</option>")
			}
			$(select).val(card.count);
			$(select).change (function(){
				var name = $(this).prop("id");
				var value = parseInt($(this).val());
				var changed = false;
				var localSettings = gameSettings.retreiveSettings();
				for (var j = 0; j < localSettings.cards.length; j++) {
					var c = localSettings.cards[j];
					if(c.card_name == name) {
						if(c.count != value) {
							localSettings.cards[j].count = value;
							changed = true;
						}
					}
				}
				if(changed) {
					gameSettings.saveSettings(localSettings);
				}
			})
		}
		view.updateNames(settings);
	},

	onGameStateChanged : function(state) {
		var game = $("#game");
		removeChildren(game);

		$(game).append("<p>Nacht: "+ state.night +"</p>").append("<p>Ronde: "+ state.round+"</p>");

		$(game).children().each (function(index, el) {
			$(el).toggleClass("col-xs-6 col-sm-3");
		});


		var players = $("#players");
		removeChildren(players);

		for (var i = 0; i < state.players.length; i++) {
			var player = state.players[i];
			var prototype = $('#players .prototype');
			var clone = $(prototype).clone().insertBefore( prototype ).toggleClass("prototype");
			$(clone).find(".player-name").html(player.id+": "+player.name);


			var cardSelect = $(clone).find(".card-select");
			var cardName = $(clone).find(".card-name");
			if(player.card.name == "unknown") {
				$(cardSelect).find("label").prop("for", "rol-"+player.id);
				var select = $(cardSelect).find("select").prop("id", "rol-"+player.id);

				var cardsAlreadyInList = [];

				for (var j = 0; j < state.unassigned.length; j++) {
					if (cardsAlreadyInList.indexOf(state.unassigned[j].name) == -1) {
						$(select).append("<option>" + state.unassigned[j].name + "</option").prop("playerId", player.id);
						cardsAlreadyInList.push(state.unassigned[j].name);
					}
				}
				$(select).change (function() {
					gameState.updateCard(parseInt($(this).prop("playerId")), $(this).val());
					repo.validateGameState(gameState.retreiveState(), gameState.saveState);
				});
				$(cardName).remove();
			} else {
				$(cardName).html(player.card.name);
				$(cardSelect).remove();
			}
		}

		var warnings = $("#warnings");
		removeChildren(warnings);
		for (var i = 0; i < state.errors.length; i++) {
			var error = state.errors[i];
			var prototype = $(warnings).find(".prototype");
			var clone = $(prototype).clone().insertBefore( prototype ).toggleClass("prototype");
			$(clone).html(error.description);
		}

	}
}

var checkLocalStorage = function() {
	var settings = gameSettings.retreiveSettings();

	if(settings != null) {
		view.onSettingsChanged(settings);
	}

	var state = gameState.retreiveState();
	if(state != null) {
		view.onGameStateChanged(state);
	}
}

var addFormListeners = function() {
	$( "#updateNamesForm" ).submit(function( event ) {
		event.preventDefault();
		var settings = gameSettings.retreiveSettings();
		settings.player_names = settings.player_names.slice(0, settings.total);
		ajax.post("new", settings, gameState.saveState);
	});
	$( "#players" ).submit(function( event ) {
		event.preventDefault();
		repo.validateGameState(gameState.retreiveState(), gameState.saveState);
	});
}

var repo = {
	getAvailableCards : function(succes) {
		ajax.get("available_cards", succes);
	},
	validateGameState : function(state, succes) {
		ajax.post("validate", state, succes);
	},
}

var ajax = {
	post : function(endpoint, data = {}, success) {
		$.ajax({
			type: "POST",
			url: "http://192.168.0.102:1237/" + endpoint,
			data : JSON.stringify(data),
			dataType: "json",
			contentType: "application/json; charset=utf-8",
			success: success,
		});
	},
	get : function(endpoint, success) {
		$.ajax({
			type: "GET",
			url: "http://192.168.0.102:1237/" + endpoint,
			dataType: "json",
			contentType: "application/json; charset=utf-8",
			success: success,
		});
	}
}

$(document).ready(function(){
	repo.getAvailableCards(gameSettings.updateAvailableCards);
	checkLocalStorage();
	addFormListeners();
});

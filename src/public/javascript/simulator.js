var newDefArray = function(n, value) {
	var array = new Array(n);
	for (var i = 0; i < array.length; i++) {
		array[i] = value;
	}
	return array;
} 

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
	}
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

		$(game).append("<p>"+ (state.night ? "Nacht" : "Dag") + " " + state.round + "</p>");
		if(state.over) {
			$(game).append("<p> De " +state.winner+ " hebben gewonnen!</p>");
		}

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
			if(!player.alive) {
				$(clone).toggleClass("alert-danger");
			} else {
				$(clone).toggleClass("alert-success");
			}
			if(player.card.name == "unknown") {
				$(cardSelect).find("label").prop("for", "rol-"+player.id);
				var select = $(cardSelect).find("select").prop("id", "rol-"+player.id);

				var cardsAlreadyInList = [];

				for (var j = 0; j < state.unassigned.length; j++) {
					if (cardsAlreadyInList.indexOf(state.unassigned[j].name) == -1) {
						$(select).append("<option>" + state.unassigned[j].name + "</option>").prop("playerId", player.id);
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
				for (var k = 0; k < player.card.targets; k++) {
					var stuff = $('<select class="form-control" ><option>-1</option></select>');
					$(stuff).prop("target", k).prop("playerId", player.id);
					for (var j = 0; j < state.players.length; j++) {
						var p = state.players[j];
						$(stuff).append("<option>" + p.id + "</option>");
					}
					if(player.actions.length > 0) {
						$(stuff).val(player.actions[0].data[k]);
					}
					$(clone).append(stuff);
				}
				var select = $(clone).find("select");
				
				$(select).change(function(){
					var value = $(this).val();
					var playerId = parseInt($(this).prop("playerId"));
					var target = parseInt($(this).prop("target"));


					var localState = gameState.retreiveState();
					var index = playerId -1;
					if(localState.players[index].actions.length == 0) {
						localState.players[index].actions = [{type:localState.players[index].card.name, data:newDefArray(localState.players[index].card.targets, "-1")}];
					}
					localState.players[index].actions[0].data[target] = value;

					gameState.saveState(localState);
					repo.validateGameState(gameState.retreiveState(), gameState.saveState);
				});

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

		//actions

		var actions = $("#actions");
		removeChildren(actions);
		var prototype = $(actions).find(".prototype");
		var clone = $(prototype).clone().insertBefore(prototype).toggleClass("prototype");
		var select = $(clone).find("select");
		for (var i = 0; i < state.players.length; i++) {
			var player = state.players[i];
			$(select).append("<option>" + player.id + "</option>");
		}
		if(state.actions.length > 0) {
			$(select).val(state.actions[0].data[0]);
		}

		$(select).change (function(){
			var state = gameState.retreiveState();
			var action = {type:"kill", data: [$(this).val()]}
			state.actions = [action];
			gameState.saveState(state);
			repo.validateGameState(gameState.retreiveState(), gameState.saveState);
		});
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
	$( "#actions" ).submit(function( event ) {
		event.preventDefault();
		var state = gameState.retreiveState();
		repo.updateGameState(state, gameState.saveState);
	});
}

var repo = {
	getAvailableCards : function(success) {
		ajax.get("available_cards", success);
	},
	validateGameState : function(state, success) {
		ajax.post("validate", state, success);
	},
	updateGameState : function(state, success) {
		ajax.post("update", state, success);
	},
}

var ajax = {
	url: "http://192.168.0.104:1237/",
	post : function(endpoint, data = {}, success) {
		$.ajax({
			type: "POST",
			url: ajax.url + endpoint,
			data : JSON.stringify(data),
			dataType: "json",
			contentType: "application/json; charset=utf-8",
			success: success,
		});
	},
	get : function(endpoint, success) {
		$.ajax({
			type: "GET",
			url: ajax.url + endpoint,
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

var hideAllCards = function() {
	$('#allCards').hide();
	$('#detail').show();
}

var showAllCards = function() {
	$('#detail').hide();
	$('#allCards').show();
}

var showCard = function(hash) {
	hideAllCards();
	$('#' + hash).show();
}

var hideCard = function(hash) {
	showAllCards();
	$('#' + hash).hide();
}

var processHash = function(hash) {
	if(hash !== "") {
		showCard(hash);
	} else {
		showAllCards();
	}
}

var getHash = function() {
	return window.location.hash.substr(1);
}

var onBack = function() {
	hideCard(getHash());
}

$(document).ready(function(){
	processHash(getHash());
});
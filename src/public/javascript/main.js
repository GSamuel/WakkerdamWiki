var hideAllCards = function() {
	$('#allCards').hide();
	$('#detail').show();
	window.scrollTo(0,0); //scroll to top
}

var showAllCards = function() {
	$('#detail').hide();
	$("#detail").children().hide(); 
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

var onHashChange = function() {
	processHash(getHash());
}

var getHash = function() {
	return window.location.hash.substr(1);
}

var onBack = function() {
	window.history.back();
}

$(document).ready(function(){
	processHash(getHash());
	window.onhashchange = onHashChange;
});
$(document).ready(function(){
	$(".card").each( function(index, element){
		if((index + 1) % 4 == 0) {
			$(element).after("<div class=\"clearfix\"></div>");
		} else if ((index+3) % 4 == 0) {
			$(element).after("<div class=\"clearfix visible-xs\"></div>");
		}
	});
});
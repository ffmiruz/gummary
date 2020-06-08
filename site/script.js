const $ = _ => document.querySelector(_)
const res = $('#results')
const lead = $('#lead')
const loader = document.createElement("div")
loader.classList.add("loader")
const Model = {
	requesting : false,
	results : "",
	input : "",
	firstQuery : true,
	receive : function (item) {
		if (item.length > 0) {
			Model.results = item
		} else {
			Model.results = "..Cannot summarize link.."
		}
		lead.removeChild(loader)
		m.redraw()
	},
}
const listResult = function (res){

	return m("div", {class:"row"},
			m("div"), 
			m("div", {class:"col 6 card"},res),
			m("div"),
		)
}
const Result = {
    view: function() {
    	if (Model.results=="") {
    		return 
    	}
return listResult(Model.results)
    }
}
$('input#term').addEventListener("keyup", function(event) {
  if (event.keyCode === 13) {
    event.preventDefault()
    $('button').click()
  }
}); 
const About = {
		view : function() {
			if (lead.parentNode){
			lead.parentNode.removeChild(lead)
			}
            return m("div",{class:"c ph2"},
            		m("h1", "ABOUT"),
            		m("a",{href:"https://github.com/ffmiyo/gummary", class:"out"}, "Source"),  
            		)
        }
	}
const Contact = {
		view : function() {
			if (lead.parentNode){
			lead.parentNode.removeChild(lead)
			}
            return m("div",{class:"c ph2"},
            		m("h1", "CONTACT"),
            		m("div", "wiccc8@gmail.com")
            		)
        }
	}
m.route(res, "/", {
	"/": Result,
	"/about": About,
	"/contact": Contact,
})
function query() {
	  Model.input = $('#term').value
	  if (!Model.input.trim()) {
	  	return //todo: input required prompt
	  }
	  if (Model.firstQuery) {
		Model.firstQuery=false
	  }
	lead.appendChild(loader)  	
	Model.results =""
	m.redraw()
	let url = "/.netlify/functions/api?q=" + Model.input
	fetch(url, {method: 'POST'})
  			.then(r => r.json())
  			.then(data => Model.receive(data.item)) 
  			.catch( _ => Model.receive( "...Error: Can't fetch results..." ))
}
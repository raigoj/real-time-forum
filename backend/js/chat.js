var form1 = document.getElementById("chatForm1");
var form2 = document.getElementById("chatForm2");

var accessMsg1 = document.getElementById('message1');
var accessMsg2 = document.getElementById('message2');


var displayMsg1 = document.getElementById('messageS1');
var displayMsg2 = document.getElementById('messageS2');

var mesValue1 = 'You: ' + accessMsg1.value;
var mesValue2 = 'You: ' + accessMsg2.value;


function handleForm(event) {
    event.preventDefault();
}
form1.addEventListener('submit', handleForm);
form2.addEventListener('submit', handleForm);

function sendMessage1() {
    displayMsg1.innerHTML += 'You: ' + accessMsg1.value + "<br>";
    displayMsg2.innerHTML += 'He: ' + accessMsg1.value + "<br>";
    scrollToBottom('messageS1');
    form1.reset();
}

function sendMessage2() {
    displayMsg1.innerHTML += 'He: ' + accessMsg2.value + "<br>";
    displayMsg2.innerHTML += 'You: ' + accessMsg2.value + "<br>";
    scrollToBottom('messageS2');

    form2.reset();
}

function scrollToBottom(id) {
    var div = document.getElementById(id);
    div.scrollTop = div.scrollHeight - div.clientHeight;
}


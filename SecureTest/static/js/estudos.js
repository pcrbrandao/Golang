// JavaScript Documentvar estudo = new Estudo("");function Estudo (elemento) {    this.elemento = elemento;}Estudo.prototype.atualizaInnerHTML = function (id, string) {    this.elemento = document.getElementById(id);    this.elemento.innerHTML = string;}Estudo.prototype.impares = function (inicio, fim) {    var texto = "";    for (i = inicio; i < fim; i++) {        if(i % 2 != 0)            texto += i + " ";    }    return texto;}Estudo.prototype.numeroParaReal = function(numero) {    var formatado = numero.toFixed(2);    formatado = "R$ " + formatado;    return formatado.replace(".", ",");}Estudo.prototype.realParaNumero = function (string) {    var compativelComParseFloat = string.replace("R$ ", "");    return parseFloat(compativelComParseFloat.replace(",", "."));}/* precisa de um input e um elemento para mostrartestaInput: function (idMensagem, idInput) {    var message, valor;    message = document.getElementById(idMensagem);    valor = document.getElementById(idInput).value;    try {        if(valor == "")  throw "empty";        if(isNaN(valor)) throw "not a number";        if(valor < 5)    throw "too low";        if(valor > 10)   throw "too high";        valor = Number(valor);        message.innerHTML = "Input is " + valor;    }    catch(err) {        message.innerHTML = "Input is " + err;    }} */// retorna um número aleatório entre mim e max incluídosfunction getRndInteger(min, max) {    return Math.floor(Math.random() * (max - min + 1) ) + min;}
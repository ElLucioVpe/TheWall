console.log("Estas re puto");

$( document ).ready(function() {
    var dateSQL = new Date().toISOString().slice(0, 19).replace('T', ' ');

    var txtDate = document.getElementById("txtDate");

    console.log(dateSQL);

    txtDate.value = dateSQL;
});
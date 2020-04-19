$(document).ready(function () {
  $("#imgInp").change(function () {
    readURL(this);
  });
});

function readURL(input) {
  if (input.files && input.files[0]) {
    var reader = new FileReader();

    reader.onload = function (e) {
      $("#imgUpload").css("background-image", "url(" + e.target.result + ")");
      console.log("aa");
      $("#imgUpload").css("background-size", "cover");
    };

    reader.readAsDataURL(input.files[0]); // convert to base64 string
  }
}

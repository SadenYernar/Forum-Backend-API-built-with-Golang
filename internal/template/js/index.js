function handleData()
{
    var form_data = new FormData(document.querySelector("form"));
    console.log(form_data.has("category"))
    if(!form_data.has("category"))
    {
        document.getElementById("chk_option_error").style.visibility = "visible";
      return false;
    }
    else
    {
        document.getElementById("chk_option_error").style.visibility = "hidden";
      return true;
    }
    
}
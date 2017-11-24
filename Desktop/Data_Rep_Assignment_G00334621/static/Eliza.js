// G00334621
// Data Representation Project 2017
// Christian Olim
// https://stackoverflow.com/questions/47425453/html-css-auto-scroll-page-to-bottom/47425655#47425655
// https://getbootstrap.com/

const form = $("#user-input");
const list = $("#conversation_list");

form.keypress(function(event){
    if(event.keyCode != 13){
        return;
    }
    event.preventDefault();
    
    // This receives the input and converts it to a value
    const userText = form.val(); 
    form.val(" ");
    
    // This makes sure the input isn't empty
    list.append("<figure><figcaption>User</figcaption></figure><li  class='list-group-item  text-left list-group-item-danger' id='leftList'>"+userText + "</li>");

    // Constructor
    const queryParams = {"user-input" : userText }
    $.get("/ask", queryParams)
    
    .done(function(resp){
        const newItem = "<figure><figcaption id ='figRight'>Eliza</figcaption></figure><li  class='list-group-item text-right list-group-item-info' id='rightList'>"+ resp + "</li>";
        setTimeout(function(){
            // Here we add a time out to make it seem more realistic
            list.append(newItem)
                $("html, body").scrollTop($("body").height());
            }, 5000);
            }).fail(function(){
            // This displays an error if the connection fails
                const newItem = "<figure><figcaption id ='figRight'>Eliza</figcaption></figure><li  class='list-group-item text-right list-group-item-danger' id='rightList'>Sorry I'm away right now!</li>";
                list.append(newItem);
            
    });
    //  This keeps the window scrolling as the conversation continues
    window.scrollTo(0, document.body.scrollHeight);
});
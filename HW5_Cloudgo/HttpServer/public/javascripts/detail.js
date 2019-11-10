window.onload = function()
{
	$.post('/loaddata/',function(data)
	{
		var store = data.split(',');
		$('#input_username').val(store[0]);
		$('#input_studentID').val(store[1]);
		$('#input_phone').val(store[2]);
		$('#input_email').val(store[3]);
	});
};

$('#log_out').click(function(){
	$('log_out').click = null;
	$.post('/logout',function(data){
		window.location.href = data;
	});
})

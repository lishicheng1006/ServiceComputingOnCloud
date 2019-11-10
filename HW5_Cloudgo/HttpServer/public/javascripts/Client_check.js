var isValid = {
	username:{
		validity:false,
		tips:"用户名6~18位英文字母、数字或下划线，必须以英文字母开头",
		estimate:function(username)
		{
			var reg = /^[a-zA-Z][a-zA-Z0-9_]{5,17}$/;
			this.validity = reg.test(username);
		}
	},
	studentID:{
		validity:false,
		tips:"学号8位数字，不能以0开头",
		estimate:function(studentID)
		{
			var reg = /^[1-9][0-9]{7}$/;
			this.validity = reg.test(studentID);
		}
	},
	password:{
		validity:false,
		tips:"密码为6~12位数字、大小写字母、中划线、下划线",
		estimate:function(password)
		{
			var reg = /^[0-9a-zA-Z-_]{6,12}$/;
			this.validity = reg.test(password);
		}
	},
	repeat_password:{
		validity:false,
		tips:"两次密码不一致",
		estimate:function(repeat_password)
		{
			this.validity = ($('#input_password').val()===repeat_password);
		}
	},
	phone:{
		validity:false,
		tips:"电话11位数字，不能以0开头",
		estimate:function(phone)
		{
			var reg = /^[1-9][0-9]{10}$/;
			this.validity = reg.test(phone);
		}
	},
	email:{
		validity:false,
		tips:"非法邮箱地址",
		estimate:function(email)
		{
			var reg = /^[a-zA-Z_\-]+@(([a-zA-Z_\-])+\.)+[a-zA-Z]{2,4}$/;
			this.validity = reg.test(email);
		}
	},
	is_checked : false,
	estimate:function()
	{
		this.is_checked = this.username.validity&&this.studentID.validity&&this.password.validity&&this.repeat_password.validity&&this.phone.validity&&this.email.validity;
	}
}

$(function(){
	$('.textinput').blur(function(){
		isValid[$(this).attr('name')].estimate($(this).val());
		if(isValid[$(this).attr('name')].validity==false)
		{
			$('#warn_'+$(this).attr('name')).css('visibility','visible');
		}
		else
		{
			$('#warn_'+$(this).attr('name')).css('visibility','hidden');
		}
	});
	window.onbeforeunload = function(){
		if($('#input_username').val()!=''||$('#input_studentID').val()!=''||$('#input_password').val()!=''||$('#input_repeat_password').val()!=''||$('#input_phone').val()!=''||$('#input_email').val()!='')
		{
			return 'warning';
		}
	};
	$('#submit_all').click(function(){
	$.ajax({
		type:"POST",
		url:"/submit/",
		async:true,
		data:$('#whole_user').serialize(),
		dataType:"text",
		beforeSend:function(){
			window.onbeforeunload = null;
			$('#submit').attr('disabled',true);
			$('.textinput').blur();
			isValid.estimate();
			if(isValid.is_checked==false)
			{
				$('#submit').attr('disabled',false);
				return false;
			}
		},
		success:function(err)
		{
			$('#submit').attr('disabled',false);
			if(err=='Yes')
			{
				window.location.href='http://localhost:3000/detail.html';
			}
			else
			{
				$('#error_show').html(err);
				window.onbeforeunload = function(){
					if($('#input_username').val()!=''||$('#input_studentID').val()!=''||$('#input_password').val()!=''||$('#input_repeat_password').val()!=''||$('#input_password').val()!=''||$('#input_repeat_password').val()!=''||$('#input_phone').val()!=''||$('#input_email').val()!='')
					{
						return 'warning';
					}
				}
			}
		}
		});
	});
});
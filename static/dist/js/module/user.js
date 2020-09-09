var User = {
	changePassword:function(){
		ui.box({
			title:'Change Password',
			size: 'mid',   //sm , lg
			backdrop:true,
			keyboard:true,
			show:true,
			remote:'/sys/user_change_password',
			method:'GET',
			btn:'<button type="button" class="btn btn-success btn-sm" onclick="$(\'#passwordForm\').submit()" id="sbmBtn">Submit</button>'
		});
	}
}
ItemSelector = {
	initialize:function(obj){
		var _this = this;
		var leftSection = $(obj.find('select')[0]);
		var rightSection = $(obj.find('select')[1]);
		var leftUpButton = $(obj.find('button')[1]);
		var leftDownButton = $(obj.find('button')[0]);
		var leftButton = $(obj.find('button')[2]);
		var rightButton = $(obj.find('button')[3]);
		var rightUpButton = $(obj.find('button')[5]);
		var rightDownButton = $(obj.find('button')[4]);
		
		_this.bindOptionDbClickEvent(leftSection,rightSection);
		_this.bindLeftButtonClickEvent(leftButton,rightButton,leftSection,rightSection);
		_this.bindMoveUpEvent(leftUpButton,leftSection);
		_this.bindMoveDownEvent(leftDownButton,leftSection);
		
		_this.bindMoveUpEvent(rightUpButton,rightSection);
		_this.bindMoveDownEvent(rightDownButton,rightSection);
	},
	bindOptionDbClickEvent(leftSection,rightSection){
		var _this = this;
		leftSection.on('dblclick','option',function(){
			var options = $(this);
			_this.appendOption(rightSection,options);
		});
		rightSection.on('dblclick','option',function(){
			var options = $(this);
			_this.appendOption(leftSection,options);
		});
	},
	bindLeftButtonClickEvent(leftButton,rightButton,leftSection,rightSection){
		var _this = this;
		leftButton.on('click',function(){
			var options = leftSection.find('option:selected');
			_this.appendOption(rightSection,options);
		});
		
		
		rightButton.on('click',function(){
			var options = rightSection.find('option:selected');
			_this.appendOption(leftSection,options);
		})
	},
	bindMoveUpEvent:function(btn,target){
		btn.click(function(){
			var selected = target.find(":selected");
		    var before = target.find(":selected:first").prev();
		    if (before.length > 0){
		    	selected.detach().insertBefore(before);
		    }
	    });
	},
	bindMoveDownEvent:function(btn,target){
		btn.click(function(){
			 var selected = target.find(":selected");
			 var next = target.find(":selected:last").next();
			 if (next.length > 0){
				 selected.detach().insertAfter(next);
			 }
		});
	},
	appendOption(target,options){
		$(options).each(function(){
			var _opt = $(this);
			var option = '<option value='+_opt.val()+'>'+_opt.text()+'</option>';
			target.append( option );
			_opt.remove();
		})
	}
}

$(function(){
	ItemSelector.initialize( $('#selectBox') );
});

var eof = -1;

exports.lexer = function(input){
  var lexer, start=0, pos=0, width=0, tokens=[];

  lexer = {}

  lexer.run = function(state){
    while(state != null){
      state = state(this);
    }
  }

  lexer.nextChar = function(){
    if(pos >= len(input)){
      width = 0;
      return eof;
    }
    var c = input[pos]
    width = 1
    pos += width
    return c
  }

  lexer.backup = function(){
    pos -= width
  }

  lexer.peek = function(){
    var c = this.nextChar();
    this.backup();
    return c;
  }

  lexer.accept = function(valid){
    if(valid.indexOf(this.nextChar()) >= 0){
      return true;
    }
    this.backup();
    return false;
  }

  lexer.acceptRun = function(valid){
    while(valid.indexOf(this.nextChar()) >= 0){}
    this.backup();
  }

  lexer.until = function(stop){
    while(true){
      var c = this.nextChar();
      if(c == eof)return;
      if(stop.indexOf(c) >= 0)break;
    }
    this.backup();
  }

  lexer.isDone = function(){
    return this.peek() == eof;
  }

  lexer.ignore = function(){
    start = pos;
  }

  lexer.emit = function(type){
    tokens.push({
      type: type,
      value: input.slice(pos, pos + width)
    });
    start = pos;
  }

  lexer.tokens = function(){
    return tokens;
  }

  return lexer;
}

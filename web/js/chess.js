function onDrop(source, target, piece, newPos, oldPos, orientation) {
  
}

var config = {
    position: 'start',
    draggable: true,
    onDrop: onDrop
  }
  
  
  var board1 = Chessboard('myBoard', config);
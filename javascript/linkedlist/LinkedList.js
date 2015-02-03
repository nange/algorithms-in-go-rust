// Node构造函数
function Node(element) {
  this.element = element;
  this.prev = null;
  this.next = null;
}

function LinkedList() {
  this.head = new Node('head');
}

LinkedList.prototype.find = function(elem) {
  var pos = this.head;
  while (pos.next !== null) {
    pos = pos.next;
    if (pos.element === elem) return pos;
  }
  return null;
};

LinkedList.prototype.findLast = function() {
  var pos = this.head;
  while (pos.next !== null) {
    pos = pos.next;
  }
  return pos;
};

LinkedList.prototype.display = function() {
  var pos = this.head;
  while (pos.next !== null) {
    console.log(pos.next.element);
    pos = pos.next;
  }
};

LinkedList.prototype.dispReverse = function() {
  var pos = this.findLast();
  while (pos.prev !== null) {
    console.log(pos.prev.element);
    pos = pos.prev;
  }
};

LinkedList.prototype.insert = function(newElem) {
  var node = new Node(newElem);
  var last = this.findLast();
  if (last === null) {
    this.head.next = node;
    node.prev = this.head;
  }

  last.next = node;
  node.prev = last;
};

LinkedList.prototype.insertAfter = function(newElem, afterElem) {
  var node = new Node(newElem);
  var pos = this.find(afterElem);

  if (pos === null) return;
  if (pos.next !== null) {
    node.next = pos.next;
    pos.next.prev = node;
  }

  node.prev = pos;
  pos.next = node;
};

LinkedList.prototype.remove = function(elem) {
  var pos = this.find(elem);
  if (pos === null) return;

  if (pos.next !== null) {
    pos.next.prev = pos.prev;
  }
  pos.prev.next = pos.next;
  pos.prev = null;
  pos.next = null;
};


module.exports = LinkedList;

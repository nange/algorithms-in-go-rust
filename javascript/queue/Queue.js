// Queue构造函数
function Queue() {
  this.dataStore = [];
}


Queue.prototype.enqueue = function(element) {
  this.dataStore.push(element);
};

Queue.prototype.dequeue = function() {
  return this.dataStore.shift();
};

Queue.prototype.size = function() {
  return this.dataStore.length;
};

Queue.prototype.isEmpty = function() {
  return this.dataStore.length == 0;
};

Queue.prototype.front = function() {
  if (this.isEmpty()) return null;
  return this.dataStore[0];
};

Queue.prototype.end = function() {
  if (this.isEmpty()) return null;
  return this.dataStore[this.size()-1];
};

Queue.prototype.clear = function() {
  this.dataStore.length = 0;
};


module.exports = Queue;

// Stack构造函数
function Stack() {
  this.dataStore = [];
}

Stack.prototype.push = function(element) {
  this.dataStore.push(element);
};

Stack.prototype.pop = function() {
  return this.dataStore.pop();
};

Stack.prototype.peek = function() {
  return this.dataStore[this.size()-1];
};

Stack.prototype.clear = function() {
  this.dataStore.length = 0;
};

Stack.prototype.size = function() {
  return this.dataStore.length;
};

Stack.prototype.isEmpty = function() {
  return this.dataStore.length === 0;
};

module.exports = Stack;

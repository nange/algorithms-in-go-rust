// 迭代器
function Iterator(list) {
  this.list = list;
  this.current = 0;
}

Iterator.prototype.hasNext = function() {
  return this.current < this.list.size();
};

Iterator.prototype.next = function() {
  if (!this.hasNext()) {
    throw new Error('No Such Element Exception!');
  }

  return this.list.get(current++);
};

Iterator.prototype.remove = function() {
  this.list.remove(--current);
};


// ArrayList 构造函数
function ArrayList() {
  this.dataStore = [];
}

ArrayList.prototype.size = function() {
  return this.dataStore.length;
};

ArrayList.prototype.get = function(index) {
  if (index < 0 || index >= this.size()) {
    throw new Error('index:' + index + ' out of bounds exception!');
  }

  return this.dataStore[index];
};

ArrayList.prototype.set = function(index, newVal) {
  if (index < 0 || index >= this.size()) {
    throw new Error('index:' + index + ' out of bounds exception!');
  }

  this.dataStore[index] = newVal;
};

ArrayList.prototype.insert = function(index, val) {
  if (index < 0 || index >= this.size()) {
    throw new Error('index:' + index + ' out of bounds exception!');
  }

  this.dataStore.splice(index, 0, val);
};

ArrayList.prototype.append = function(val) {
  this.dataStore.push(val);
};

ArrayList.prototype.remove = function(index) {
  if (index < 0 || index >= this.size()) {
    throw new Error('index:' + index + ' out of bounds exception!');
  }

  this.dataStore.splice(index, 1);
};

ArrayList.prototype.clear = function() {
  this.dataStore = [];
};

ArrayList.prototype.find = function(val) {
  for (var i = 0; i < this.dataStore.length; i++) {
    if (this.dataStore[i] === val) return i;
  }
  return -1;
};

ArrayList.prototype.iterator = function() {
  return new Iterator(this);
};

ArrayList.prototype.toString = function() {
  return this.dataStore.join();
};

module.exports = ArrayList;

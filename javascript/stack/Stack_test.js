var assert = require('assert');
var Stack = require('./Stack.js');

describe('Stack', function() {
  describe('#size(), #push()', function() {
    it('should return 1 when push a element into Stack.', function() {
      var stack = new Stack();
      stack.push('1');
      assert.equal(1, stack.size());
    });
  });

  describe('#pop(), #peek()', function() {
    it('should return valid value and length.', function() {
      var stack = new Stack();
      stack.push('1');
      stack.push('2');
      stack.push('3');
      stack.push('4');

      assert.equal(4, stack.pop());
      assert.equal(3, stack.size());
      assert.equal('3', stack.peek());
    });
  });

  describe('#clear()', function() {
    it('should return 0, when Stack called clear()', function() {
      var stack = new Stack();
      stack.push('1');
      stack.push('2');
      stack.push('3');
      stack.clear();
      assert.equal(0, stack.size());
    });
  });

});

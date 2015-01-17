var assert = require('assert');
var Queue = require('./Queue.js');

describe('Queue', function() {
  describe('#enqueue, #dequeue, #size, #isEmpty, #front, #end, #clear', function() {
    it('should all success.', function() {
      var q = new Queue();
      q.enqueue(1);
      q.enqueue(2);
      q.enqueue(3);

      assert.equal(3, q.size());
      assert.equal(false, q.isEmpty());

      var d = q.dequeue();
      assert.equal(1, d);
      assert.equal(2, q.size());
      assert.equal(3, q.end());
      assert.equal(2, q.front());

      q.clear();
      assert.equal(true, q.isEmpty());

    });
  });

});

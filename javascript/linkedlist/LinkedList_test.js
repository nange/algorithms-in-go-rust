var assert = require('assert');
var LinkedList = require('./LinkedList.js');

describe('Queue', function() {
  describe('#insert, #insertAfter, #find, #findLast, #remove', function() {
    it('should all success.', function() {
      var list = new LinkedList();
      list.insert('1');
      list.insert('2');
      list.insert('3');
      assert.equal('2', list.find('3').prev.element);
      assert.equal('3', list.findLast().element);

      list.insertAfter('4', '2');
      assert.equal('4', list.find('3').prev.element);
      assert.equal('4', list.find('2').next.element);
      assert.equal('3', list.findLast().element);

      list.remove('4');
      assert.equal('2', list.find('3').prev.element);

    });
  });

});

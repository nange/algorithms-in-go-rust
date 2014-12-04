var assert = require('assert');
var ArrayList = require('../lib/ArrayList.js');

describe('ArrayList', function() {
  describe('#size()', function() {
    it('should return 0 when the ArrayList is empty', function() {
      var arrayList = new ArrayList();
      assert.equal(0, arrayList.size());
    });

    it('should return valid value', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);
      arrayList.append(3);
      assert.equal(3, arrayList.size());

      arrayList.remove(2);
      assert.equal(2, arrayList.size());
    });
  });

  describe('#get(index)', function() {
    it('should return valid value', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);
      arrayList.append(3);
      assert.equal(3, arrayList.get(2));
      assert.equal(2, arrayList.get(1));
    });
  });

  describe('#set(index, newVal)', function() {
    it('should set success and have no error', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);
      arrayList.append(3);
      assert.equal(3, arrayList.get(2));

      assert.doesNotThrow(function() {
        arrayList.set(2, 4);
      });
      assert.equal(4, arrayList.get(2));

      assert.throws(function() {
        arrayList.set(3, 4);
      }, Error);
    });
  });

  describe('#insert(index, val)', function() {
    it('should insert success and have no error', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      assert.doesNotThrow(function() {
        arrayList.insert(0, 2);
      });
      
      assert.equal(2, arrayList.get(0));
      assert.equal(2, arrayList.size());
    });
  });

  describe('#append(val)', function() {
    it('should append success and have no error', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      
      assert.equal(1, arrayList.get(0));
      assert.equal(1, arrayList.size());
    });
  });

  describe('#remove(index)', function() {
    it('should remove success and have no error', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);
      arrayList.append(3);
      assert.equal(3, arrayList.size());

      arrayList.remove(0);
      assert.equal(2, arrayList.get(0));
      assert.equal(2, arrayList.size());
    });
  });

  describe('#clear()', function() {
    it('should clear success and have no error', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);
      arrayList.append(3);

      arrayList.clear();
      assert.equal(0, arrayList.size());
    });
  });

  describe('#find(val)', function() {
    it('should return a valid value', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);
      arrayList.append(3);

      arrayList.find(2);
      assert.equal(1, arrayList.find(2));
    });
  });

  describe('#iterator()', function() {
    it('should return a valid value', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);
      var iterator = arrayList.iterator();
      assert.notEqual(undefined, iterator);
      assert.notEqual(undefined, iterator.next);
    });
  });

  describe('#toString()', function() {
    it('should return a valid value', function() {
      var arrayList = new ArrayList();
      arrayList.append(1);
      arrayList.append(2);

      assert.equal('1,2', arrayList.toString());
    });
  });

});

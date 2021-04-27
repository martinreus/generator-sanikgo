var sanitize = require("../sanitize")
var fs = require('fs');
const SuperGenerator = require('../super-generator');

module.exports = class extends SuperGenerator {

  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);

  }

  initializing() {
    this.templateConfig = {
    }

  }

  async prompting() {
  }

  configuring() { }


  async writing() {
    await this._copyFiles("", "", this.templateConfig)
  }

  install() { }
  end() {
  }

};

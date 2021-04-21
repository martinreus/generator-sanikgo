var Generator = require('yeoman-generator');
var sanitize = require("../sanitize")

module.exports = class extends Generator {

  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);

  }

  initializing() {
    this.templateConfig = {
      moduleName: sanitize.appName(this.appname) // Default to current folder name
    }

    this.log("app name", this.templateConfig.moduleName);
  }

  async prompting() {
  }

  configuring() { }


  writing() { }

  install() { }
  end() {
    console.log("finished base app generation");
  }

};

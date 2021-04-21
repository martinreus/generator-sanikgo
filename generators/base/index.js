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
  }

  async prompting() {
  }

  configuring() { }


  writing() {
    this.fs.copyTpl(this.templatePath(`go.mod`),
      this.destinationPath(`go.mod`),
      this.templateConfig)
  }

  install() { }
  end() {
    console.log("finished base app generation");
  }

};
